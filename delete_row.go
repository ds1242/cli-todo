package main

import (
	"strconv"
	"errors"
	"os"
)


func deleteRecord(recordId string) error {
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

	// remove element from the slice
	updatedSlice := RemoveElement(records, recordIdInt)

	// delete data file to clear it out
	deleteErr := os.Remove("data.csv")
	if deleteErr != nil {
		return errors.New("error replacing csv file")
	}	
	
	for _, slice := range updatedSlice {
		writeErr := writeToCSV(slice)
		if writeErr != nil {
			return writeErr
		}
	}
	// exit delete
	return nil
}

// this func removes and element and shifts the rest of the elements to the right keeping the order
func RemoveElement(s [][]string, index int) [][]string {
	return append(s[:index], s[index+1:]...)
}