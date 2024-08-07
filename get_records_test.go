package main

import (
	"fmt"
	"testing"
)

func TestGetRecords(t *testing.T) {
	records := GetRecords("data.csv", ",")
	fmt.Println(records)
}