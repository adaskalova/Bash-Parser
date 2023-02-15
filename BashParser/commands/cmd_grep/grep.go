package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	colPurple = "\033[35m"
	colNone   = "\033[0m"
	colRed    = "\033[0;31m"
)

func changeColor(s string) string {
	return colPurple + s + colNone
}

func changeColRed(s string) string {
	return colRed + s
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
		return "The value you enter isn't valid! Please, enter a valid command!", err
	}

	return
}

func readOutput(reader io.Reader, prefix string) (string, error) {
	rdr := bufio.NewReader(reader)
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
	return "", nil
}

func verifyOutExecuteCmd(input string) (string, error) {
	out, err := executeCmd(input)
	if out == "" {
		return "", nil
	}
	if err != nil {
		fmt.Println(".....")
	}
	return "", nil
}

func main() {

	args := os.Args
	input := args[1]

	verifyOutExecuteCmd(input)

}
