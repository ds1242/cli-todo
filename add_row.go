package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func addRow(value string) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return errors.New("unable to open data.csv")
	}

	filedata, _ := csv.NewReader(file).ReadAll()

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	id := strconv.Itoa(len(filedata))
	createdAt := time.Now().UTC
	completed := "false"

	rowToAdd := []string{id, value, createdAt().String(), completed }
	if err := writer.Write(rowToAdd); err != nil {
		
		return errors.New("error writing record to file")
	}

	return nil
}