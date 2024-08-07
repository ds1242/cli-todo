package main

import (
	// "encoding/csv"
	"fmt"
	"strconv"
	// "strings"
	"errors"
)


func deleteRecord(recordId string)(string, error) {
	recordIdInt, recordIdErr:= strconv.Atoi(recordId)
	if recordIdErr != nil {
		return "", errors.New("unable to parse record id")
	}
	var recordStructSlice []Row
	records := GetRecords("data.csv", ",")
	fmt.Println(recordIdInt)
	// fmt.Println(records)
	for _, record := range records {
		recordStructSlice = append(recordStructSlice, convertSliceOfRowsToRowStructs(record))
	}
	fmt.Println(recordStructSlice)
	return "Deleted", nil
}