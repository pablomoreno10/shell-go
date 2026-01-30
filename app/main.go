package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var _ = fmt.Print

func main() {

	exit := "exit"
	echo := "echo"
	builtin := "type"
	pwd := "pwd"

	for {
		fmt.Print("$ ")

		//read user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		//invalid input
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		//Clean and Tokenize ONCE
		cleanInput := strings.TrimSpace(input)
		parts := strings.Fields(cleanInput)

		// Handle empty input (user just hit Enter)
		if len(parts) == 0 {
			continue
		}

		//Separate Command from Arguments
		cmd := parts[0]
		args := parts[1:]

		//Commands
		if cmd == builtin {
			//Logic for 'type'
			target := strings.Join(args, " ")

			if target == exit || target == builtin || target == echo || target == pwd {
				fmt.Println(target + " is a shell builtin")
			} else {
				path, err := exec.LookPath(target)
				if err != nil {
					fmt.Println(target + ": not found")
				} else {
					fmt.Println(target + " is " + path)
				}
			}

		} else if cmd == echo {
			//Logic for 'echo'
			output := strings.Join(args, " ")
			fmt.Println(output)

		} else if cmd == exit {
			//Logic for 'exit'
			os.Exit(0)

		} else if cmd == pwd {
			//Logic for 'pwd'
			command, err := os.Getwd()
			if err != nil {
				fmt.Println("Error running command: %v\n", err)
			} else {
				fmt.Println(command)
			}
		
		}	else {
			//Logic for external executables
			_, err := exec.LookPath(cmd)
			if err != nil {
				fmt.Println(cmd + ": command not found")
			} else {
				//We use the 'cmd' variable and the 'args' slice we created at the top
				command := exec.Command(cmd, args...) //'...' unpacks the args slice
				
				//Standard Output + Standard Error combined
				output, err := command.CombinedOutput()
				if err != nil {
					fmt.Printf("Error running command: %v\n", err)
				} else {
					fmt.Print(string(output))
				}
			}
		}
	}
}
