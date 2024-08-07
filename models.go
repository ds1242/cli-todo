package main

import (
	"log"
	"strconv"
	"time"
)

type Row struct {
	ID 			int
	Task		string
	CreatedAt	time.Time
	Completed	bool
}

func convertSliceOfRowsToRowStructs(record []string) Row {
	idInt, idErr := strconv.Atoi(record[0])
	if idErr != nil {
		log.Fatal("unable to parse ID")
		return Row{}
	}

	// List of potential layouts
	layouts := []string{
		time.RFC3339,
		time.UnixDate,
		time.RFC1123Z,
		// Add custom layouts as needed
	}
	var timeClean time.Time
	var err error

	for _, layout := range layouts {
		timeClean, err = time.Parse(layout, record[2])
		if err == nil {
			break
		}
	}
	
	completedBool, boolErr := strconv.ParseBool(record[3])
	if boolErr != nil {
		log.Fatal("unable to parse bool value")
		return Row{}
	}

	return Row{
		ID: 		idInt,
		Task: 		record[1],
		CreatedAt: 	timeClean,
		Completed: 	completedBool,
	}

}