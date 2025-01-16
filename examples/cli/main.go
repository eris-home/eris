package main

import (
	"fmt"
	"github.com/eris-home/eris/examples/cli"
)

func main() {
	fmt.Println("\nWelcome to the Expanded Eris CLI Example!")
	fmt.Println("=========================================")

	runner := cli.NewRunner()
	runner.Start()
}