package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	commandList := map[string]string{"add": "add", "delete": "delete", "list": "list", "complete": "complete", "help":"help"}

	userArgs := os.Args
	if len(userArgs) < 3 && (userArgs[1] != "list" && userArgs[1] != "help") {		
		log.Fatal("not enough arguments")
	}
	
	if len(userArgs) > 3 {
		log.Fatal("too many arguments")
	}

	command, errCheckCommand := checkCommand(userArgs[1], commandList)
	if errCheckCommand != nil {
		log.Fatalf("Error parsing args: %v\n", errCheckCommand)
	}

	if command == "add" {
		addErr := addRow(userArgs[2])
		if addErr != nil {
			fmt.Printf("Error adding row: %v\n", addErr)
			return
		}
	}

	if command == "delete" {
		err := deleteRecord(userArgs[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Record Deleted")
	}

	if command == "list" {
		err := listTasks()
		if err != nil {
			fmt.Printf("error listing tasks: %v", err)
			return 
		}
	}
	
	if command == "complete" {
		err := completeRecord(userArgs[2])
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	if command == "help" {
		fmt.Print(`A todo list for the terminal

Usage: 
	tasks [commmand]

Available Commands:
	add			[task]	Add a new task
	complete 	[id]	Mark the task complete with the provided id
	delete		[id]	Delete the task with the provided id
	list				List all of the tasks in the todo list
	help				Help about any command


	`)
	}

}

