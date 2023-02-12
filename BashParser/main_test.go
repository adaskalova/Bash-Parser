package main

import (
	"testing"
)

var path = "commands/test_files/test/rm_file.txt"
var changePath = "cd /tmp"

func BenchmarkIsPathExists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPathExists(path)
	}
}

func BenchmarkChangeDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		changeDir(changePath)
	}
}
