package main

import (
	"os"
	"encoding/csv"
	"errors"
)
func addRow(val string) error {
	file, err := os.Open("data.csv")
	if err != nil {
		return errors.New("unable to open data.csv")
	}

	defer file.Close()
}