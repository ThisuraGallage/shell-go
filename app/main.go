package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			return
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			return
		}

		if strings.HasPrefix(command, "echo ") {
            args := command[5:]
            fmt.Println(args)
            continue
        }


		fmt.Println(command + ": command not found")

	}
}
