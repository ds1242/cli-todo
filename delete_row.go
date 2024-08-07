package main

import (
	// "encoding/csv"
	// "fmt"
	"strconv"
	// "strings"
	"os"
	"errors"
)


func deleteRecord(recordId string)(string, error) {
	// convert input from user to get the record id as an int
	recordIdInt, recordIdErr:= strconv.Atoi(recordId)
	if recordIdErr != nil {
		return "", errors.New("unable to parse record id")
	}

	mapOfRows := make(map[int]Row)
	records := GetRecords("data.csv", ",")

	// Load map with row objects
	for i, record := range records {
		mapOfRows[i] = convertSliceOfRowsToRowStructs(record)
	}
	
	for _, row := range mapOfRows {
		if row.ID == recordIdInt {
			delete(mapOfRows, row.ID)
		}
	}

	deleteErr := os.Remove("data.csv")
	if deleteErr != nil {
		return "", errors.New("error replacing csv file")
	}	

	file, err := os.Create("data.csv")
	if err != nil {
		return "", errors.New("error creating a new file")
	}
	file.Close()

	for _, row := range mapOfRows {
		addErr := addRow(row.Task)
		if addErr != nil {
			return "", errors.New("error updating the task list after deletion")
		}
	}
	
	
	return "Record Deleted", nil
}