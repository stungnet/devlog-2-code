package main // Main package that will create executable

import (
	"fmt" // Format and print functions for outputting data to console.
	"os" // Operating system functions.
)

// Main function is entry point for the program, where code begins execution when executed.
func main (){

	// filename will store user input.
	var filename string
	// Ask for user input.
	fmt.Printf("Name the txt filename contents you want to read: ")
	// Store user input.
	fmt.Scan(&filename)
	filename += ".txt" // add .txt extension to user input.

	// Here we take the filename given with the added .txt and look for that file to read.
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err) // If file doesn't exist it will throw error.
		os.Exit(1) // Exit program with error 1.
	}

	// If filename exists we take that data and place it into a content variable then print it to the user.
	content := string(data)
	fmt.Println(content)
}
