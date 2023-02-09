package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	colPurple       = "\033[35m"
	colNone         = "\033[0m"
	colRed          = "\033[0;31m"
	pipeEscaped     = `(?:[^\\|]|\\[\s\S])+`
	pathExists      = `[\/]`
	matchWhiteSpace = `[\/]\s+`
)

var (
	bFlag      bool
	bHasPrefix bool
	bPath      bool
	bSpace     bool
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

func isPipeEscaped(str string) bool {
	result := regexp.MustCompile(pipeEscaped).MatchString(str)
	return result
}

func isPathExists(path string) bool {
	result := regexp.MustCompile(pathExists).MatchString(path)
	return result
}
func isMatchWhiteSp(str string) bool {
	result := regexp.MustCompile(matchWhiteSpace).MatchString(str)
	return result
}

func changeDir(str string) (msg string) {
	bSpace = isMatchWhiteSp(str)
	if bSpace {
		return
	}
	words := strings.Fields(str)
	wordsLen := len(words)
	var tmpMsg string
	if wordsLen > 1 {
		tmpWords := words[1:]
		for _, item := range tmpWords {
			bPath := isPathExists(item)
			if bPath {
				err := os.Chdir(filepath.Join("", item))
				//os.Getwd(): return (dir string, err error)
				//	string:   current directory
				//	error:    if any
				tmpMsg, err = os.Getwd()
				fmt.Println("tmpMsg:", tmpMsg)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return tmpMsg
}

func executeCmd(inputCmd string) (s string, err error) {

	cmd := exec.Command("bash", "-c", inputCmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return
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
	if err != nil {
		fmt.Println(".....")
	}

	return "", nil
}

func readOutput(reader io.Reader, prefix string) {
	rdr := bufio.NewReader(reader)
	// result := ""
	bs := []byte{}
	for {
		bs, _, _ = rdr.ReadLine()
		if bs != nil {
			outStr := string(bs)
			coloredTxt := changeColor(outStr)
			fmt.Println(prefix + coloredTxt)
		} else {
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

		output = isPipeEscaped(input)
		if output {

			input = strings.ReplaceAll(input, "'", ".")
			bHasPrefix = strings.HasPrefix(input, "cd")

			if bHasPrefix {
				msg := changeDir(input)
				tmp := isEmpty(msg)
				if !tmp {
					coloredMsg := changeColor(msg)
					fmt.Println("Curren working directory:", coloredMsg)
				} else {
					fmt.Println("The value you enter isn't valid! Please, enter a valid command!")
				}
			}
			executeCmd(input)
		}

		if input == "Q" || input == "q" {
			// break from the for loop
			break
		}
	}

}
