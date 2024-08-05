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
		log.Fatalf("Error parsing args: %v", errCheckCommand)
	}


	fmt.Printf("Args: %v\n", command)
	fmt.Printf("Command list: %v\n", commandList)

}

