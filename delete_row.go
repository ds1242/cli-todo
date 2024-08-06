package main

import (
	"encoding/csv"
	"fmt"
	"strings"
	"errors"
)


func deleteRecord(recordId string)(string, error) {
	records, err := csv.NewReader(strings.NewReader("data.csv")).ReadAll()
	if err != nil {
		return "", errors.New("unable to delete record")
	}
	// for _, record := range records {
	// 	if record[0] == recordId {

	// 	}
	// }
	fmt.Println(recordId)
	fmt.Println(records)
	return "", nil
}