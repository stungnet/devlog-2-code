package main //Main package to that will create executable

// Imports from standard library
import (
	"fmt" // Format and print functions for outputting data to console.
	"os" // Operating system functions.
	"strings" // String manipulation functions.
)

// Main function is our entry point for the program where our code begins execution when executed.
func main(){
	var input string // Our string which user input will be stored.
	
	// The beginning of our for loop
	for {
		// Ask user question with fmt.Printf, then take the input with fmt.Scanln.
		fmt.Printf("Please enter 'a' or 'b': ")
		fmt.Scanln(&input)

		// If statement will check that user has not enter just blank. This will only check the first entry not subsequent.
		if len(strings.TrimSpace(input)) == 0{
			fmt.Println("You didn't enter anything? ¯(°_o)/¯")
			os.Exit(1)
		}

		// We take that use input and convert it to lowercase.
		userInputLower := strings.ToLower(input)

		// Let the user know what they entered.
		fmt.Println("You entered", userInputLower)

		// If block will check to see if "a" or "b" if neither user will be asked again until it is "a" or "b".
		if userInputLower == "a"{
			fmt.Println("You listened!")	
			break
		} else if userInputLower == "b"{
			fmt.Println("Why did you go b??")
			break
		} else{
			fmt.Println("Input incorrect try again.")
		}
	}
}
