package main

import (
	"errors"
	"strings"
)

func parseAndCleanArgs(args string, commandList map[string]string) (string, error) {
	if len(args) < 2 {		
		return "", errors.New("not enough arguments")
	}
	
	if len(args) > 3 {
		return "", errors.New("too many arguments")
	}

	val, ok := commandList[strings.ToLower(args)] 
	if !ok {
		return "", errors.New("invalid command entered")
	}

	return val, nil
}