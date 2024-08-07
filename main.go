package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	commandList := map[string]string{"add": "add", "delete": "delete", "list": "list"}

	userArgs := os.Args
	if len(userArgs) < 3 {		
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
		message, err := deleteRecord(userArgs[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(message)
	}
	
	fmt.Printf("Args: %v\n", command)
	fmt.Printf("Value: %s\n", userArgs[2])
	fmt.Printf("Type of: %T\n", userArgs[2])

}

