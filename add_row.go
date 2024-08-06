package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"bufio"
	"time"
	"fmt"
)

func addRow(value string) error {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("unable to open data.csv")
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	id, err := generateId("data.csv")
	if err != nil {
		return errors.New("error getting length of file")
	}

	createdAt := time.Now().UTC
	completed := "false"

	rowToAdd := []string{id, value, createdAt().String(), completed }
	
	writeErr := writer.Write(rowToAdd)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func generateId(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("unable to open data.csv")
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// getting the number of lines
	filedata, err := csv.NewReader(file).ReadAll()
    if err != nil {        
        return "", err
    }

	lineCount := len(filedata)
	fmt.Println("Number of lines in file:", lineCount)
	idVal := lineCount + 1

	id:= strconv.Itoa(idVal)

	return id, nil
}

