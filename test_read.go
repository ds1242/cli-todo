package main


import "fmt"

func ExampleMain() {
	records := GetRecords("data.csv", ",")
	fmt.Println(records)
}

