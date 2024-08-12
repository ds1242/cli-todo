package main

import (
	"errors"
	"fmt"
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

	fmt.Println(records)

	return nil
}