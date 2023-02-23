package main

import (
	"fmt"
	"os"
)


func main (){

	var filename string
	fmt.Printf("Name the txt filename contents you want to read: ")
	fmt.Scan(&filename)
	filename += ".txt" // add .txt extension to user input.

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // give error code 1 to be known by.
	}

	content := string(data)
	fmt.Println(content)
}
