package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {

	builtins := []string{"echo", "exit", "type", "pwd"}

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

		if command == "pwd" {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
			} else {
				fmt.Println(dir)
			}
			continue
		}

		if strings.HasPrefix(command, "echo ") {
			args := command[5:]
			fmt.Println(args)
			continue
		}

		if strings.HasPrefix(command, "type ") {
			args := strings.TrimSpace(command[5:])

			isBuiltin := false
			for _, builtin := range builtins {
				if args == builtin {
					isBuiltin = true
					break
				}
			}

			if isBuiltin {
				fmt.Printf("%s is a shell builtin\n", args)
				continue
			}

			pathEnv := os.Getenv("PATH")
			pathDirs := strings.Split(pathEnv, string(os.PathListSeparator))

			found := false
			for _, dir := range pathDirs {
				fullPath := filepath.Join(dir, args)

				// Check if file exists
				fileInfo, err := os.Stat(fullPath)
				if err != nil {
					continue
				}

				// not a directory
				if !fileInfo.Mode().IsRegular() {
					continue
				}

				// Check execute permissions
				if fileInfo.Mode()&0111 != 0 {
					fmt.Printf("%s is %s\n", args, fullPath)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("%s: not found\n", args)
			}
			continue
		}

		// Handle external programs
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue
		}

		cmdName := parts[0]
		cmdArgs := parts[1:]

		// Check if it's a builtin (skip builtins)
		isBuiltin := false
		for _, builtin := range builtins {
			if cmdName == builtin {
				isBuiltin = true
				break
			}
		}

		if !isBuiltin {
			// Use exec.Command directly with cmdName - it will search PATH automatically
			cmd := exec.Command(cmdName, cmdArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			err := cmd.Run()
			if err != nil {
				fmt.Println(command + ": command not found")
			}
			continue
		}

		fmt.Println(command + ": command not found")

	}
}
