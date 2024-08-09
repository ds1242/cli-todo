package main

import (
	"encoding/csv"
	"os"
	"errors"
)



func writeToCSV(row []string) error {
	if !fileExists("data.csv") {
		os.Create("data.csv")
	}
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("unable to open data.csv")
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	writeErr := writer.Write(row)
	if writeErr != nil {
		return writeErr
	}
	return nil
}