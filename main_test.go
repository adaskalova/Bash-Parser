package main

import (
	"testing"
)

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

func TestIsPipeEscaped(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "isPipeEscaped", input: "asWG |4mcv | 7rT2J|IFVkBjv", expected: true},
		{desc: "isPipeEscaped", input: "asWG|4mcv|7rT2J|IFVkBjv", expected: true},
		{desc: "isPipeEscaped", input: "asWG4mcv7rT2JIFVkBjv", expected: true},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := isPipeEscaped(tc.input)
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
		{desc: "isPathExists", input: "/", expected: true},
		{desc: "isPathExists", input: "path/to/the/file.txt", expected: true},
		{desc: "isPathExists", input: "At4QHT1M8nIucRlugXzt", expected: false},
		{desc: "isPathExists", input: `path\to\the\file.txt`, expected: false},
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

func TestIsMatchWhiteSp(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "isMatchisMatchWhiteSp", input: "/ abc", expected: true},
		{desc: "isMatchisMatchWhiteSp", input: "/abc", expected: false},
		{desc: "isMatchisMatchWhiteSp", input: "/", expected: false},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output := isMatchisMatchWhiteSp(tc.input)
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
		{desc: "changeDir", input: "cd /", expected: "/"},
		{desc: "changeDir", input: "cd /tmp", expected: "/tmp"},
		{desc: "changeDir", input: `cd \`, expected: ""},
		{desc: "changeDir", input: `cd \tmp`, expected: ""},
		{desc: "changeDir", input: "cd / tmp", expected: ""},
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

func TestExecuteCmd(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "executeCmd", input: "cd /", expected: ""},
		{desc: "executeCmd", input: "cd", expected: ""},
		{desc: "executeCmd", input: "cd /t mp", expected: "/tmp"},
		{desc: "executeCmd", input: "cd /tmp", expected: "/tmp"},
		{desc: "executeCmd", input: "cd df", expected: "/tmp"},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			output, err := executeCmd(tc.input)
			if err != nil {
				t.Fatalf("output:  %v; expected:  %v", output, tc.expected)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
