package main

import (
	"fmt"
)


func listTasks() error {
	tasks := GetRecords("data.csv", ",")

	fmt.Println("===============================================================================================")
	fmt.Println("| ID  |        Task           |           Created At               |   Status  |")
	count := 0
	for _, task := range tasks {
		// for _, col := range task {
			fmt.Println("-----------------------------------------------------------------------")
			fmt.Printf("|  %d  |     %s      |   %s   |   %s  |\n", count, task[0], task[1], task[2])
			// fmt.Printf(col[0])
			count++
		// }
		
	}
	fmt.Println("===============================================================================================")

	return nil
}