package main // Main package that will create executable

import (
	"bufio" // Buffered input and output opertations; Improve reading and writing.
	"fmt" // Format and print functions for outputting data to console.
	"os" // Operating system functions
)

// Main function is entry point for the program, where code begins execution when executed.
func main(){
	var filename string // String to store new filename
	fmt.Printf("Give your new .txt file a name: ") // Prompt user.
	fmt.Scan(&filename)	// Take user input.

	// Here we add .txt to the users inputted filename.
	filename += ".txt" 
	
	// Creating filename.txt
	os.Create(filename)
	fmt.Printf("%v created!", filename) // Let user know it has being created.

	// Prompt user to enter text to our new filename.txt
	fmt.Println("What do you want to write to the file? ")
	// Create a new bufio.Reader that reads from os.Stdin, which is our users input.
	scanner := bufio.NewReader(os.Stdin) 
	// Using the `scanner` we created just before we ReadString until the end of the line and assign it to contents.
	contents, _ := scanner.ReadString('\n')
	
	// Convert the input to byes to write it to the file.
	data := []byte(contents)
	os.WriteFile(filename, data, 0664)
}
