package main

import (
	"fmt"
	"testing"
)

// TODO: Mock up test better to actually be useful
func TestGetRecords(t *testing.T) {
	records := GetRecords("data.csv", ",")
	fmt.Println(records)
}