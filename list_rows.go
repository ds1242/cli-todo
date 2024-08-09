package main

import (
	"fmt"
)


func listTasks() error {
	tasks := GetRecords("data.csv", ",")
	fmt.Println(tasks)

	return nil
}