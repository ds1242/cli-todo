package main

import (
	"fmt"
)


func listTasks() error {
	tasks := GetRecords("data.csv", ",")

	fmt.Println("==================================================")
	fmt.Println("| ID |     Task      |   Created At  |   Status  |")
	for _, task := range tasks {
		for _, col := range task {
			fmt.Println("--------------------------------------------------")
			fmt.Printf("|  |     %v      |   %v |   %v  |\n", col[0], col[1], col[2])
			
		}
		
	}
	fmt.Println("==================================================")

	return nil
}