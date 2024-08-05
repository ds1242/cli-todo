package main

import (
	"errors"
	"strings"
)

func checkCommand(args string, commandList map[string]string) (string, error) {
	
	val, ok := commandList[strings.ToLower(args)] 
	if !ok {
		return "", errors.New("invalid command entered")
	}

	return val, nil
}