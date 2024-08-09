package main

import (
	"time"
)

func addRow(task string) error {
	
	createdAt := time.Now().UTC()
	completed := "false"

	rowToAdd := []string{ task, createdAt.String(), completed }
	
	err := writeToCSV(rowToAdd)
	if err != nil {
		return err
	}

	return nil
}


