package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintData(s string) {
	for _, char := range s {
		z01.PrintRune((char))
	}
}

func ReadFile(fName string) string {
	info, err := ioutil.ReadFile(fName)
	if err != nil {
		return "error"
	}
	return string(info)
}

func main() {
	arg := os.Args[1:]
	bool := false
	for _, files := range arg {
		if _, err := os.Stat(files); err != nil {
			PrintData("ERROR: open " + files + ": no such file or directory\n")
			os.Exit(1)
		}
		PrintData(ReadFile(files))
		bool = true
	}
	if !bool {
		ioutil.ReadAll(io.TeeReader(os.Stdin, os.Stdout))
		os.Stdin.Close()
		os.Stdout.Close()
	}
}
