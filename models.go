package main

import (
	"strconv"
	"time"
	"os"
)

type Row struct {
	Task		string
	CreatedAt	time.Time
	Completed	bool
}


func convertRowStructToSlice(record Row) []string {

	completedString := strconv.FormatBool(record.Completed)
	
	rowToAdd := []string{record.Task, record.CreatedAt.String(), completedString}

	return rowToAdd
}


func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
	   return false
	}
	return !info.IsDir()
 }