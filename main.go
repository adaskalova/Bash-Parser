package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

var (
	bFlag  bool
	bDigit bool
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

func isNum(str string) bool {
	for _, ch := range str {
		if unicode.IsNumber(ch) {
			fmt.Println(string(ch), "is a number.")
			bDigit = true
		} else {
			//is not a number rune
			bDigit = false
		}
	}
	return bDigit
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

		outEmpty := isEmpty(input)
		if outEmpty {
			fmt.Println("The value you enter isn't valid! Please, enter a valid command!")
		}

		outDigit := isNum(input)
		if outDigit {
			fmt.Println("The value you enter isn't valid! Please, enter a valid command!")
		}

		if input == "Q" || input == "q" {
			// break from the for loop
			break
		}
	}

}
