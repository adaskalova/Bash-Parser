package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	bFlag     bool
	bAlphaNum bool
)

func isEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		// if string is empty
		bFlag = true
	} else {
		bFlag = false
	}
	return bFlag
}

func isPipeSep(str string) bool {
	result := regexp.MustCompile(`[^0-9](?:[^\\|]|\\[\s\S])+`).MatchString(str)
	if result {
		// is an alphanumeric
		bAlphaNum = true
	} else {
		//is not an alphanumeric
		bAlphaNum = false
	}
	return bAlphaNum
}

func main() {
	var input string
	//path to the program
	name := os.Args[0]
	//return the file name
	name = path.Base(name)
	//display the program name
	fmt.Println("Program Name: " + name)
	//read the input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your command:")
		//scans the input
		scanner.Scan()
		//get string typed in the standard input
		input = scanner.Text()

		output := isEmpty(input)
		if output {
			fmt.Println("The value you enter isn't valid! Please, enter a valid command!")
		}

		output = isPipeSep(input)
		if output {

			//TODO

		}

		if input == "Q" || input == "q" {
			// break from the for loop
			break
		}
	}

}
