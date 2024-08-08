package main

import (
	// "encoding/csv"
	"fmt"
	"strconv"
	// "strings"
	"errors"
	// "os"
)


func deleteRecord(recordId string) error {
	// convert input from user to get the record id as an int
	recordIdInt, recordIdErr:= strconv.Atoi(recordId)
	if recordIdErr != nil {
		return errors.New("unable to parse record id")
	}
	// var sliceOfRows [][]string
	records := GetRecords("data.csv", ",")

	// Load map with row objects
	// for i, record := range records {
	// 	// TODO: not parsing date correctly
	// 	sliceOfRows[i] = convertSliceOfRowsToRowStructs(record)
	// }
	
	if len(records) == 0 {
		return errors.New("no row to delete")
	}

	// for _, row := range sliceOfRows {
	// 	if row[0] == recordIdInt {
	// 		delete(sliceOfRows, row.ID)
	// 	}
	// }
	updatedSlice := RemoveElement(records, recordIdInt)
	// deleteErr := os.Remove("data.csv")
	// if deleteErr != nil {
	// 	return errors.New("error replacing csv file")
	// }	

	// _, err := os.Create("data.csv")
	// if err != nil {
	// 	return errors.New("error creating a new file")
	// }

	// file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return errors.New("unable to open data.csv")
	// }
	
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	// for _, row := range sliceOfRows {
	// 	rowStringSlice := convertRowStructToSlice(row)
	// 	writeErr := writer.Write(rowStringSlice)
	// 	if writeErr != nil {
	// 		return writeErr
	// 	}
	// }
	
	
	return nil
}

func RemoveElement(s [][]string, index int) [][]string {
	return append(s[:index], s[index+1:]...)
}