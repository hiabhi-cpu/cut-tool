package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {

		check(commandEntry(os.Args))

	}

}

func commandEntry(argsCmd []string) error {
	isCsv := true
	hasD := false
	hasF := false
	hasFile := false
	dFlag := "\t"
	fileName := ""
	var fList []int

	for i, rec := range argsCmd {
		if strings.HasPrefix(rec, "-f") {
			filedArgs := ""
			if rec == "-f" && i+1 < len(argsCmd) {
				filedArgs = argsCmd[i+1]
			} else {
				filedArgs = strings.Trim(argsCmd[i], "-f")
			}
			fList = getFiledNum(filedArgs)

			hasF = true

		}
		if strings.HasPrefix(rec, "-d") {
			hasD = true
			dFlag = strings.Trim(argsCmd[i], "-d")
			if dFlag == "" {
				check(errors.New("Give a delimiter for -d"))

			}
		}
		if strings.Contains(rec, ".csv") || strings.Contains(rec, ".tsv") {
			hasFile = true
			if strings.Contains(rec, ".csv") {
				isCsv = true
			} else {
				isCsv = false
			}
			fileName = argsCmd[i]
		}
	}

	var reader *csv.Reader

	if !hasFile {

		data, err := readFromConsole()
		check(err)
		reader = csv.NewReader(strings.NewReader(string(data)))
		if strings.Contains(string(data), "\t") {
			isCsv = false
		}
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		reader = csv.NewReader(file)
	}
	if !hasF {
		check(errors.New("Give an -f paramter"))
	}

	if isCsv {
		commandCSV(reader, fList, hasD, dFlag)
	} else {
		commandTSV(reader, fList, hasD, dFlag)
	}

	return nil
}

func commandCSV(reader *csv.Reader, field []int, hasD bool, dFlag string) error {
	if len(field) == 0 {
		return errors.New("Give correct column number > 0")
	}
	for {
		record, err := reader.Read()
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return err
		}
		if !hasD {
			fmt.Println(record, ",")
			continue
		}
		if dFlag != "," {
			fmt.Println(record, ",")
			continue
		}
		for _, r := range field {
			fmt.Print(record[r-1], ",")
		}
		fmt.Println()

	}

	return nil
}

func commandTSV(reader *csv.Reader, field []int, hasD bool, dFlag string) error {

	if len(field) == 0 {
		return errors.New("Give correct column number > 0")
	}
	for {
		record, err := reader.Read()
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return err
		}
		if dFlag != "\t" {
			fmt.Println(record, "\t")
			continue
		}
		row := strings.Split(record[0], "\t")
		for i, rec := range row {
			for _, r := range field {
				if i+1 == r {
					fmt.Print(rec, "\t")
				}
			}
		}
		fmt.Println()
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFiledNum(str string) []int {

	temp := make([]string, 0)
	if strings.Contains(str, " ") {
		temp = strings.Split(str, " ")
	} else {
		temp = strings.Split(str, ",")
	}
	listField := make([]int, len(temp))
	for i, r := range temp {
		tempInt, err := strconv.Atoi(r)
		check(err)
		listField[i] = tempInt
	}
	return listField
}

func readFromConsole() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	return io.ReadAll(reader)
}
