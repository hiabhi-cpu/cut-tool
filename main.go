package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) >= 1 {
		if strings.Contains(os.Args[len(os.Args)-1], ".csv") || strings.Contains(os.Args[len(os.Args)-1], ".tsv") { //has file name

			// fmt.Println(argsCmd)
			// fmt.Println(fileName)
			check(commandEntry(os.Args))
		} else {
			return
		}
	}

}

func commandEntry(argsCmd []string) error {
	isCsv := true
	hasD := false
	hasF := false
	// fFlag := 0
	dFlag := "\t"
	fileName := ""
	var fList []int
	// fmt.Println(argsCmd)
	// for _, rec := range argsCmd {
	// 	fmt.Println(rec)
	// }
	// return nil
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
			// fmt.Println(dFlag)
		}
		if strings.Contains(rec, ".csv") || strings.Contains(rec, ".tsv") {
			if strings.Contains(rec, ".csv") {
				isCsv = true
			} else {
				isCsv = false
			}
			fileName = argsCmd[i]
		}
	}
	if !hasF {
		check(errors.New("Give an -f paramter"))
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	if isCsv {
		commandCSV(reader, fList, hasD, dFlag)
	} else {
		commandTSV(reader, fList, hasD, dFlag)
	}

	if err != nil {
		return err
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
