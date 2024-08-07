package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func addRow(task string) error {
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
	
	id, err := checkFileLength("data.csv")
	if err != nil {
		return errors.New("error getting length of file")
	}

	idVal := id + 1
	idString:= strconv.Itoa(idVal)

	createdAt := time.Now().UTC()
	completed := "false"

	rowToAdd := []string{idString, task, createdAt.String(), completed }
	
	writeErr := writer.Write(rowToAdd)
	if writeErr != nil {
		return writeErr
	}

	return nil
}



