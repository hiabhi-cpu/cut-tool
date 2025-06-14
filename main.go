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
	fFlag := 0
	dFlag := "\t"
	fileName := ""
	for i, rec := range argsCmd {
		if strings.HasPrefix(rec, "-f") {
			fieldString, err := strconv.Atoi(strings.Trim(argsCmd[i], "-f"))
			hasF = true
			check(err)
			fFlag = fieldString
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
		commandCSV(reader, fFlag, hasD, dFlag)
	} else {
		commandTSV(reader, fFlag, hasD, dFlag)
	}

	if err != nil {
		return err
	}
	fmt.Println(hasD)
	return nil
}

func commandCSV(reader *csv.Reader, field int, hasD bool, dFlag string) error {
	if field == 0 {
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
			fmt.Println(record)
			continue
		}
		if dFlag != "," {
			fmt.Println(record)
			continue
		}
		fmt.Println(record[field-1])
	}

	return nil
}

func commandTSV(reader *csv.Reader, field int, hasD bool, dFlag string) error {

	if field == 0 {
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
			fmt.Println(record)
			continue
		}
		row := strings.Split(record[0], "\t")
		for i, rec := range row {
			if i+1 == field {
				fmt.Println(rec)
			}
		}
		// fmt.Println(record)
		// fmt.Println(len(record))
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
