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

	builtins := []string{"echo", "exit", "type"}

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

		if strings.HasPrefix(command, "type ") {
			args := strings.TrimSpace(command[5:])
			// Check if it's a builtin command
			isBuiltin := false
			for _, builtin := range builtins {
				if args == builtin {
					isBuiltin = true
					break
				}
			}

			if isBuiltin {
				fmt.Printf("%s is a shell builtin\n", args)
			} else {
				fmt.Printf("%s: not found\n", args)
			}
			continue
		}

		fmt.Println(command + ": command not found")

	}
}
