package main

import (
	"encoding/csv"
	"io"
	"os"
	"log"
)

// TODO: adjust return to add error handling
func GetRecords(filePath, separator string) [][]string {
	records := [][]string{}
	if len(filePath) == 0 {
		return records
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return records
	}
	
	csvReader := csv.NewReader(file)
	csvReader.Comma = []rune(separator)[0]
	for {
		columnValues, err := csvReader.Read()
		if err != nil {
			break
	}
	if err == io.EOF {
		break
	}
	
	records = append(records, columnValues)
	}
	return records
}