package main

import (
	"ffxiv/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
