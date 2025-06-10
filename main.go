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
			fileName := os.Args[len(os.Args)-1]
			argsCmd := os.Args[1 : len(os.Args)-1]
			// fmt.Println(argsCmd)
			// fmt.Println(fileName)
			check(commandEntry(fileName, argsCmd))
		} else {
			return
		}
	}

}

func commandEntry(fileName string, argsCmd []string) error {
	isCsv := true
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if !strings.Contains(fileName, ".csv") {
		isCsv = false
	}

	for _, rec := range argsCmd {
		if strings.HasPrefix(rec, "-f") {
			err = commandField(reader, argsCmd[0], isCsv)
		}
	}

	if err != nil {
		return err
	}
	return nil
}

func commandField(reader *csv.Reader, cmd string, isCsv bool) error {
	fieldString := strings.Trim(cmd, "-f")
	field, err := strconv.Atoi(fieldString)
	if err != nil {
		return err
	}
	if field == 0 {
		return errors.New("Give correct column number > 0")
	}
	// fmt.Println(field)
	for {
		record, err := reader.Read()
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return err
		}
		// fmt.Println(record)
		var row []string
		if isCsv {
			row = strings.Split(record[0], " ")
		} else {
			row = strings.Split(record[0], "\t")
		}

		for i, rec := range row {
			if i+1 == field {
				fmt.Println(rec)
			}
		}
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
