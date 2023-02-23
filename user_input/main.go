package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string
	for {
		fmt.Printf("Please enter 'a' or 'b': ")
		fmt.Scan(&input)
		userInputLower := strings.ToLower(input)

		fmt.Println("You entered", userInputLower)

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
