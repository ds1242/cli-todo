package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	commandList := map[string]string{"add": "add", "delete": "delete", "list": "list"}

	userArgs := os.Args
	
	cleanArgs, errCleaningArgs := parseAndCleanArgs(userArgs[1], commandList)
	if errCleaningArgs != nil {
		log.Fatalf("Error parsing args: %v", errCleaningArgs)
	}


	fmt.Printf("Args: %v\n", cleanArgs)
	fmt.Printf("Command list: %v\n", commandList)

}

