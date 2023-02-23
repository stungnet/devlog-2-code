package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var filename string
	fmt.Printf("Give your new .txt file a name: ")
	fmt.Scan(&filename)	
	filename += ".txt" // add .txt extension to user input.

	os.Create(filename)
	fmt.Printf("%v created!", filename)

	fmt.Println("What do you want to write to the file? ")
	scanner := bufio.NewReader(os.Stdin)
	contents, _ := scanner.ReadString('\n')
	
	data := []byte(contents)
	os.WriteFile(filename, data, 0664)
}
