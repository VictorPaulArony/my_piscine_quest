package main

import (
	"fmt"
	"os"

	"lem-in/utils"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run . [FILE]")
		return
	}

	fileName := args[1]
	utils.Printout(fileName)
}
