package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	colPurple       = "\033[35m"
	colNone         = "\033[0m"
	colRed          = "\033[0;31m"
	pathExists      = `[\/]`
	matchWhiteSpace = `[\/]\s+`
)

var (
	cmtOutput string
	fInfo     bool
	bSpace    bool
	bFlag     bool
)

//change the color of the output text
func changeColor(s string) string {
	return colPurple + s + colNone
}

//run the executable file
func executeCommand(path string, arg string) (result string) {
	exe, err := exec.LookPath(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	output, err := exec.Command(exe, arg).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
	return string(output)

}

//check whether the string contains slash
func isPathExists(path string) bool {
	result := regexp.MustCompile(pathExists).MatchString(path)
	return result
}

//check for whitespace in file path
func isMatchWhiteSpace(str string) bool {
	result := regexp.MustCompile(matchWhiteSpace).MatchString(str)
	return result
}

//change the current working directory to the user-entered directory
func changeDir(str string) (msg string) {
	bSpace = isMatchWhiteSpace(str)
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
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return tmpMsg
}

// check for empty string
func isEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		// if string is empty
		bFlag = true
	} else {
		bFlag = false
	}
	return bFlag
}

//checks if the command begins with a specified prefix
//and executes it
func selectCmd(cmd string) {
	baseTmp := isContainStr(cmd, "base64")
	baseDecode := isContainStr(cmd, "base64 --decode")
	switch {
	case strings.HasPrefix(cmd, "cat"):
		cmtOutput = executeCommand("commands/cmd_cat/cat", cmd)
	case strings.HasPrefix(cmd, "sort"):
		cmtOutput = executeCommand("commands/cmd_sort/sort", cmd)
	case strings.HasPrefix(cmd, "grep"):
		cmtOutput = executeCommand("commands/cmd_grep/grep", cmd)
	case strings.HasPrefix(cmd, "mv"):
		cmtOutput = executeCommand("commands/cmd_mv/mv", cmd)
	case strings.HasPrefix(cmd, "ls"):
		cmtOutput = executeCommand("commands/cmd_ls/ls", cmd)
	case strings.HasPrefix(cmd, "tail"):
		cmtOutput = executeCommand("commands/cmd_tail/tail", cmd)
	case strings.HasPrefix(cmd, "cd"):
		cmtOutput = changeDir(cmd)
		tmp := isEmpty(cmtOutput)
		if !tmp {
			coloredMsg := changeColor(cmtOutput)
			fmt.Println("Curren working directory:", coloredMsg)
		} else {
			fmt.Println("The value you enter isn't valid! Please, enter a valid command!")
		}
	case baseTmp:
		cmtOutput = executeCommand("commands/cmd_base64/encode", cmd)
		//create temp file
		fileName := "log.txt"
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = f.WriteString(cmtOutput)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		err = f.Sync()
		if err != nil {
			panic(err)
		}
	case baseDecode:
		cmtOutput = executeCommand("commands/cmd_base64_decode/base64_decode", cmd)
		//remove file
		err := os.Remove("log.txt")
		if err != nil {
			fmt.Println(err)
		}
	case strings.HasPrefix(cmd, "rm"):
		cmtOutput = executeCommand("commands/cmd_rm/rm", cmd)
	default:
		fmt.Println("The guess is wrong!")
	}
}

// check if file exists
func checkFileExists(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//check if the string is present or not in the command
func isContainStr(str string, spec string) bool {
	output := strings.ContainsAny(str, spec)
	return output
}

//get file path
func getPath(str []string) string {
	for _, value := range str {
		fInfo = checkFileExists(value)
		if fInfo {
			return value
		}

	}
	return ""
}

//constructing a command
func constructCommand(sl string, m map[int]string) string {
	//find the length of map
	var result int = len(m)
	words := strings.Fields(sl)
	outDashes := isContainStr(sl, "--")
	outSlashes := isContainStr(sl, "/")
	outDot := isContainStr(sl, ".")

	if len(words) > 1 && outDashes {
		return sl + " " + "log.txt"
	} else if (len(words) == 1) && (result == 1) {
		output := sl + " " + m[0]
		return output
	} else if len(words) > 1 && outSlashes {
		return sl

	} else if len(words) > 1 && outDot {
		output := sl + " " + m[0]
		return output

	} else if len(words) == 2 {
		output := sl + " " + m[0]
		return output

	}
	return ""

}

//command parsing
func parseCommand(cmd string) {
	var output string
	cmds := make(map[string]string)
	strCmd := make(map[int]string)
	strReplace := strings.Replace(cmd, "'", ".", -1)
	v := strings.Split(strReplace, "|")
	for i, val := range v {
		cmds[val] = strconv.Itoa(i)
		words := strings.Fields(string(val))
		path := getPath(words)
		if path != "" {
			strCmd[i] = path
		}
		output = constructCommand(val, strCmd)
		colorOutput := changeColor(output)
		fmt.Println("Command Used: ", colorOutput)

		selectCmd(output)
	}
}

//take the user input from the command-line
func execute() {
	var cmd string

	flag.StringVar(&cmd, "cmd", "", "Please, specify command.")
	flag.Parse()
	fmt.Println(cmd)
	if len(cmd) == 0 {
		fmt.Println("Usage: ProgramName -cmd")
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		parseCommand(cmd)
	}

}

func main() {

	execute()
}
