package main

import (
	// "encoding/csv"
	"fmt"
	"strconv"
	// "strings"
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
			return "Deleted record", nil
		}
	}

	fmt.Println(mapOfRows)

	return "no record to delete", nil
}