package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	colPurple = "\033[35m"
	colNone   = "\033[0m"
	colRed    = "\033[0;31m"
)

var (
	bFlag      bool
	bSep       bool
	bHasPrefix bool
	bPath      bool
)

func changeColor(s string) string {
	return colPurple + s + colNone
}
func changeColRed(s string) string {
	return colRed + s
}

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
		bSep = true
	} else {
		bSep = false
	}
	return bSep
}

func isPathExists(path string) bool {
	match := regexp.MustCompile(`^\/[a-zA-Z]+`).MatchString(path)
	tmpMatch := regexp.MustCompile(`\/`).MatchString(path)
	if match || tmpMatch {
		return true
	} else {
		return false
	}
}

func changeDir(str string) {
	words := strings.Fields(str)
	wordsLen := len(words)
	if wordsLen > 1 {
		tmp_words := words[1:]
		for _, item := range tmp_words {
			bPath := isPathExists(item)
			if bPath {
				cwd, err := os.Getwd()
				if err != nil {
					fmt.Printf("err: %T, %v\n", err, err)
				}
				err = os.Chdir(filepath.Join("", item))
				cwd, _ = os.Getwd()
				fmt.Println("cwd:", cwd)
				if err != nil {
					log.Printf("error: %v\n", err)
				}
			}
		}
	} else {
		str := "The value you enter isn't valid! Please, enter a valid command!"
		coloredTxt := changeColor(str)
		fmt.Println(coloredTxt)
	}
}

func executeCmd(inputCmd string) error {

	cmd := exec.Command("bash", "-c", inputCmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return err
	}

	coloredTxt := changeColRed("Error: ")

	go readOutput(stdout, "Output: ")
	go readOutput(stderr, coloredTxt)

	// start the command
	err = cmd.Start()
	if err != nil {
		panic("Error: The command could not be executed!")
	}

	//waiting for command to finish
	err = cmd.Wait()
	log.Printf("Completed successfully...")

	return nil
}

func readOutput(reader io.Reader, prefix string) {
	rdr := bufio.NewReader(reader)
	result := ""
	bs := []byte{}
	for bs != nil {
		bs, _, _ = rdr.ReadLine()
		result = string(bs)
		coloredTxt := changeColor(result)
		fmt.Println(prefix + coloredTxt)
		if bs == nil {
			break
		}
	}
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
		fmt.Print("> Enter your command:")
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

			input = strings.ReplaceAll(input, "'", ".")
			bHasPrefix = strings.HasPrefix(input, "cd")

			if bHasPrefix {
				changeDir(input)
			}
			executeCmd(input)
		}

		if input == "Q" || input == "q" {
			// break from the for loop
			break
		}
	}

}
