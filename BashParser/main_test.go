package main

import (
	"testing"
)

var path = "commands/test_files/test/rm_file.txt"

func BenchmarkIsPathExists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPathExists(path)
	}
}

func BenchmarkChangeDir(b *testing.B) {
	benchmarks := []struct {
		name       string
		changePath string
	}{
		{"Path_root:", "/"},
		{"Path_tmp:", "/tmp"},
		//{"Path_random:", "<path>/<to>/<the>/<specific>/<directory>"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				changeDir(bm.changePath)
			}
		})
	}
}
