package main

import (
	"encoding/csv"
	"os"
	"bufio"
	"errors"
)


func checkFileLength(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, errors.New("unable to open data.csv")
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// getting the number of lines
	filedata, err := csv.NewReader(file).ReadAll()
    if err != nil {        
        return 0, err
    }

	lineCount := len(filedata)
	

	return lineCount, nil
}