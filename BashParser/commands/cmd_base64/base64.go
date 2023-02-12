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

var (
	bFlag      bool
	bHasPrefix bool
	bPath      bool
	bSpace     bool
)

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

	go readOutput(stdout, "")
	go readOutput(stderr, "")

	// start the command
	err = cmd.Start()
	if err != nil {
		panic("Error: The command could not be executed!")
	}

	//waiting for command to finish
	err = cmd.Wait()
	if err != nil {
		panic(err)
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
			fmt.Println(outStr)
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
		panic(err)
	}

}
