package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkChangeColRed(b *testing.B) {
	str := "output"
	for i := 0; i < b.N; i++ {
		changeColRed(str)
	}
}

func BenchmarkChangeColor(b *testing.B) {
	str := "error"
	for i := 0; i < b.N; i++ {
		changeColor(str)
	}
}

func BenchmarkExecuteCmd(b *testing.B) {
	str := "cat ../../commands/test_files/misc.txt"
	for i := 0; i < b.N; i++ {
		executeCmd(str)
	}
}

func BenchmarkVerifyOutExecuteCmd(b *testing.B) {
	benchmarks := []struct {
		desc string
		cmd  string
	}{
		{"Command:", "cat ../../commands/test_files/computer.txt"},
		{"Command_empty_string:", ""},
		{"Command_not_found:", "abc"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				verifyOutExecuteCmd(bm.cmd)
			}
		})
	}
}

func BenchmarkReadOutput(b *testing.B) {
	str := `This is a test string.This is a test string.
	This is a test string.This is a test string.`
	rdr := strings.NewReader(str)
	for i := 0; i < b.N; i++ {
		readOutput(rdr, "Output")
	}
}

func TestExecuteCmd(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "ExecuteCmd cat:", input: "cat ../../commands/test_files/computer.txt", expected: ""},
		{desc: "ExecuteCmd:", input: "cat commands/test_files/computer.txt", expected: "The value you enter isn't valid! Please, enter a valid command!"},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output, err := executeCmd(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
			if err != nil {
				fmt.Fprintln(os.Stderr)
				return
			}
		})
	}
}

func TestReadOutput(t *testing.T) {
	str := "This is a test string."
	expected := ""
	rdr := strings.NewReader(str)
	output, err := readOutput(rdr, "Output:")
	if output != expected {
		t.Fatalf("output: %s expected: %s", output, expected)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr)
		return
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

func TestVerifyOutExecuteCmd(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "Verify command:", input: "cat ../../commands/test_files/misc.txt", expected: ""},
		{desc: "Verify command:", input: "cat commands/test_files/computer.txt", expected: "The value you enter isn't valid! Please, enter a valid command!"},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output, err := verifyOutExecuteCmd(tc.input)
			if output != tc.expected {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
			if err != nil {
				fmt.Fprintln(os.Stderr)
				return
			}
		})
	}
}
