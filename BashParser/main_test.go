package main

import (
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
