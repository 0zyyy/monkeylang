package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Read the monkey program from file
	content, err := os.ReadFile("main.monk")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Executing program from main.monk\n\n\n")
	// Pass the file content directly to REPL
	repl.Start(string(content), os.Stdout)
}
