package main

import (
	"fmt"
	"os"
)

//docker run <image> <cmd> <parameters>
//go run main.go <cmd> <parameters>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
}
