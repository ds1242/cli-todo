package main

import (
	"errors"
	"os"
	"strconv"
)


func completeRecord(recordId string) error {
	// convert input from user to get the record id as an int
	recordIdInt, recordIdErr:= strconv.Atoi(recordId)
	if recordIdErr != nil {
		return errors.New("unable to parse record id")
	}

	records := GetRecords("data.csv", ",")

	// if length of records = 0 there are no records
	if len(records) == 0 {
		return errors.New("no row to delete")
	}
	// check for out of bounds index
	if recordIdInt >= len(records) {
		return errors.New("that record does not exists")
	}
	records[recordIdInt][2] = "true"
	
	// delete data file to clear it out
	deleteErr := os.Remove("data.csv")
	if deleteErr != nil {
		return errors.New("error replacing csv file")
	}	

	for _, row := range records {
		writeErr := writeToCSV(row)
		if writeErr != nil {
			return writeErr
		}
	}

	return nil
}