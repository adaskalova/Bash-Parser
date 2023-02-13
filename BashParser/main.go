package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

func changeColor(s string) string {
	return colPurple + s + colNone
}

var (
	cmtOutput string
	fInfo     bool
	bSpace    bool
	bFlag     bool
)

func readOutput(reader io.Reader, prefix string) {
	rdr := bufio.NewReader(reader)
	// result := ""
	bs := []byte{}
	for {
		bs, _, _ = rdr.ReadLine()
		if bs != nil {
			outStr := string(bs)
			fmt.Println(prefix + outStr)
		} else {
			break
		}
	}
}

func run(path string, arg string) (result string) {

	exe, _ := exec.LookPath(path)

	output, err := exec.Command(exe, arg).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
	return string(output)

}
func isPathExists(path string) bool {
	result := regexp.MustCompile(pathExists).MatchString(path)
	return result
}

func isMatchWhiteSpace(str string) bool {
	result := regexp.MustCompile(matchWhiteSpace).MatchString(str)
	return result
}

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

func isEmpty(str string) bool {
	if strings.TrimSpace(str) == "" {
		// if string is empty
		bFlag = true
	} else {
		bFlag = false
	}
	return bFlag
}

func selectCmd(cmd string) {
	baseTmp := isContainStr(cmd, "base64")
	baseDecode := isContainStr(cmd, "base64 --decode")
	switch {
	case strings.HasPrefix(cmd, "cat"):
		cmtOutput = run("commands/cmd_cat/cat", cmd)
	case strings.HasPrefix(cmd, "sort"):
		cmtOutput = run("commands/cmd_sort/sort", cmd)
	case strings.HasPrefix(cmd, "grep"):
		cmtOutput = run("commands/cmd_grep/grep", cmd)
	case strings.HasPrefix(cmd, "mv"):
		cmtOutput = run("commands/cmd_mv/mv", cmd)
	case strings.HasPrefix(cmd, "ls"):
		cmtOutput = run("commands/cmd_ls/ls", cmd)
	case strings.HasPrefix(cmd, "tail"):
		cmtOutput = run("commands/cmd_tail/tail", cmd)
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
		cmtOutput = run("commands/cmd_base64/encode", cmd)
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
		cmtOutput = run("commands/cmd_base64_decode/base64_decode", cmd)
		//remove file
		err := os.Remove("log.txt")
		if err != nil {
			fmt.Println(err)
		}
	case strings.HasPrefix(cmd, "rm"):
		cmtOutput = run("commands/cmd_rm/rm", cmd)
	default:
		fmt.Println("The guess is wrong!")
	}
}

// check file exists
func checkFileExists(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isContainStr(str string, spec string) bool {
	output := strings.ContainsAny(str, spec)
	return output
}

func getPath(str []string) string {
	for _, value := range str {
		// fmt.Println("value = ", value)
		fInfo = checkFileExists(value)
		// fmt.Println("fInfo = ", fInfo)
		if fInfo {
			return value
		}

	}
	return ""
}

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

func cmdOptions(cmd string) {
	var output string
	cmds := make(map[string]string)
	strCmd := make(map[int]string)
	strReplace := strings.Replace(cmd, "'", ".", -1)
	v := strings.Split(strReplace, "|")
	// fmt.Println("value1 = ", v)
	for i, val := range v {
		// fmt.Println("value2 = ", val)
		cmds[val] = strconv.Itoa(i)
		// fmt.Println(cmds)
		words := strings.Fields(string(val))
		// fmt.Println("words = ", words)
		path := getPath(words)
		// fmt.Println("key: ", key)
		// fmt.Println("path: ", path)
		if path != "" {
			strCmd[i] = path
		}
		// fmt.Println("cmds: ", strCmd)
		// fmt.Println("value2 = ", val)
		output = constructCommand(val, strCmd)
		colorOutput := changeColor(output)
		fmt.Println("Command Used: ", colorOutput)

		selectCmd(output)
	}
}

func main() {

	var cmd string

	flag.StringVar(&cmd, "cmd", "cat commands/default/default.txt", "Specify command. Default is cat commands/default/default.txt")
	flag.Parse()

	cmdOptions(cmd)
}
