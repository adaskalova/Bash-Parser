package main

import (
	"fmt"
	"testing"
)

func BenchmarkIsPathExists(b *testing.B) {
	var path = "commands/test_files/test/rm_file.txt"
	for i := 0; i < b.N; i++ {
		isPathExists(path)
	}
}

func BenchmarkIsContainStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isContainStr("cat commands/test_files/misc.txt | base64 | base64 --decode", "/")
	}
}

func BenchmarkIsMatchWhiteSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatchWhiteSpace("cd / tmp")
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	benchmarks := []struct {
		desc       string
		changePath string
	}{
		{"Empty string:", ""},
		{"NOT empty string:", "adf43D"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isEmpty(bm.changePath)
			}
		})
	}
}

func BenchmarkChangeDir(b *testing.B) {
	benchmarks := []struct {
		desc       string
		changePath string
	}{
		{"Path_root:", "cd /"},
		{"Path_tmp:", "cd /tmp"},
		{"Path_root:", ""},
		//{"Path_random:", "<path>/<to>/<the>/<specific>/<directory>"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				changeDir(bm.changePath)
			}
		})
	}
}

func BenchmarkSelectCmd(b *testing.B) {
	benchmarks := []struct {
		desc    string
		command string
	}{
		{"Command_cat:", "cat commands/test_files/computer.txt"},
		{"Command_sort:", "sort commands/test_files/misc.txt"},
		{"Command_grep:", "grep an commands/test_files/*.*"},
		{"Command_ls:", "ls commands/test_files"},
		{"Command_tail:", "tail commands/test_files/misc.txt"},
		{"Command_cd:", "cd /"},
		{"Command_cd:", "cd /tmp"},
		{"Command_mv:", "mv commands/test_files/misc.txt commands/test_files/test"},
		{"Command_mv:", "mv commands/test_files/test/misc.txt commands/test_files"},
		{"Command_rm:", "rm commands/test_files/test/rm_file.txt"},
		{"Command_base64 --decode:", "base64 --decode"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				selectCmd(bm.command)
			}
		})
	}
}

func BenchmarkConstructCommand(b *testing.B) {
	dict := map[int]string{0: "commands/test_files/misc.txt"}
	str := "cat commands/test_files/misc.txt | base64 | base64 --decode"
	for i := 0; i < b.N; i++ {
		constructCommand(str, dict)
	}
}

func BenchmarkGetPath(b *testing.B) {
	str := []string{"grep Apple.s commands/test_files/*.*"}
	for i := 0; i < b.N; i++ {
		getPath(str)
	}
}

func BenchmarkParseCommand(b *testing.B) {
	benchmarks := []struct {
		desc string
		str  string
	}{
		{"Command_cat:", "cat commands/test_files/computer.txt"},
		{"Command_complex:", "cat commands/test_files/misc.txt | sort | grep 'nux' | tail"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parseCommand(bm.str)
			}
		})
	}
}

func BenchmarkCheckFileExists(b *testing.B) {
	benchmarks := []struct {
		desc string
		str  string
	}{
		{"File name:", "grep"},
		{"File name:", "commands/test_files/*.*"},
		{"File name:", "commands/test_files/computer.txt"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkFileExists(bm.str)
			}
		})
	}
}

func BenchmarkExecuteCommand(b *testing.B) {
	benchmarks := []struct {
		desc string
		path string
		arg  string
	}{
		{"Execute command:", "commands/test_files", "ls -la commands/test_files"},
		{"Execute command:", "commands/test_files/misc.txt", "cat commands/test_files/misc.txt"},
		{"Execute command:", "commands/cmd_cat/cat", "cat commands/test_files/misc.txt | base64 | base64 --decode"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				executeCommand(bm.path, bm.arg)
			}
		})
	}
}

func BenchmarkChangeColor(b *testing.B) {
	str := "cat commands/test_files/misc.txt"
	for i := 0; i < b.N; i++ {
		changeColor(str)
	}
}

func TestIsMatchWhiteSpace(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "isMatchWhiteSpace_a", input: "/ abc", expected: true},
		{desc: "isMatchWhiteSpace_b", input: "/abc", expected: false},
		{desc: "isMatchWhiteSpace_c", input: "/", expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := isMatchWhiteSpace(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestChangeDir(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "changeDir_a", input: "cd /", expected: "/"},
		{desc: "changeDir_b", input: "cd /tmp", expected: "/tmp"},
		{desc: "changeDir_c", input: `cd \`, expected: ""},
		{desc: "changeDir_d", input: `cd \tmp`, expected: ""},
		{desc: "changeDir_e", input: "cd / tmp", expected: ""},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := changeDir(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		inputArg string
		expected string
	}{
		{desc: "executeCmd_a", input: "commands/cmd_cat/cat", inputArg: "cat commands/test_files/abc.txt", expected: "abcdefg"},
		{desc: "executeCmd_b", input: "commands/cmd_sort/sort", inputArg: "sort commands/test_files/abc.txt", expected: "abcdefg"},
		{desc: "executeCmd_d", input: "commands/cmd_tail/tail", inputArg: "tail commands/test_files/abc.txt", expected: "abcdefg"},
		{desc: "executeCmd_e", input: "commands/cmd_ls/ls", inputArg: "ls commands/test_files/", expected: "Output: abc.txt\nOutput: computer.txt\nOutput: misc.txt\nOutput: monitor.txt\nOutput: test\nOutput: wiki_apple.txt"},
		{desc: "executeCmd_f", input: "commands/cmd_cd/mv", inputArg: "mv commands/test_files/abc.txt commands/test_files/test", expected: "....."},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := executeCommand(tc.input, tc.input_arg)
			fmt.Println(output)
			ch := output != tc.expected
			if !ch {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}

}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		//test description
		desc string
		//function input
		input string
		//expected output
		expected bool
	}{
		// Non Empty string
		{desc: "TestIsEmpty", input: "P4bXBgNDrD", expected: false},
		{desc: "TestIsEmpty", input: " \t\n P4bXBgNDrD \n\t\r\n", expected: false},
		{desc: "TestIsEmpty", input: "$$  $$", expected: false},
		{desc: "TestIsEmpty", input: "0", expected: false},
		// Empty string
		{desc: "TestIsEmpty", input: "", expected: true},
		// string with whitespace
		{desc: "TestIsEmpty", input: "    ", expected: true},
		{desc: "TestIsEmpty", input: " \t\n  \n\t\r\n", expected: true},
	}

	for _, tc := range tests {
		//result for each test case
		t.Run(tc.desc, func(t *testing.T) {
			output := isEmpty(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestIsPathExists(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "isPathExists_a", input: "/", expected: true},
		{desc: "isPathExists_b", input: "path/to/the/file.txt", expected: true},
		{desc: "isPathExists_c", input: "At4QHT1M8nIucRlugXzt", expected: false},
		{desc: "isPathExists_d", input: `path\to\the\file.txt`, expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := isPathExists(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestChangeColor(t *testing.T) {
	str := "This is a test string."
	expected := "\033[35mThis is a test string.\033[0m"
	output := changeColor(str)
	if output != expected {
		t.Fatalf("output: %v expected: %v", output, expected)
	}

}
