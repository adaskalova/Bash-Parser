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
	fmt.Println(inputCmd)

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

	args := os.Args
	input := args[1]

	out, err := executeCmd(input)
	if out == "" {
		return
	}
	if err != nil {
		fmt.Println(".....")
	}

}
