package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/google/shlex"
)

var _ = fmt.Print

func main() {

	exit := "exit"
	echo := "echo"
	builtin := "type"
	pwd := "pwd"
	cd := "cd"
	quit := "q"

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

		// Handle empty input (user just hit Enter)
		if cleanInput == "" {
			continue
		}

		//Separate Command from Arguments
		content, _ := shlex.Split(cleanInput)
		cmd := content[0]
		args := content[1:]

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
			content:= strings.Join(args, " ")
			fmt.Println(content)
			
		} else if cmd == exit || cmd == quit{
			//Logic for 'exit'
			os.Exit(0)

		} else if cmd == pwd {
			//Logic for 'pwd'
			command, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error running command: %v\n", err)
			} else {
				fmt.Println(command)
			}
		
		} else if cmd == cd {
			//Logic for 'cd'
			path := strings.Join(args, " ")
			if path == "~"{
				home, err := os.UserHomeDir()
				if err != nil{
					fmt.Printf("Error reaching home directory: %v\n", err)
				}
				if os.Chdir(home) != nil{
					fmt.Println("Error reaching home directory")
				}
			}else{
				if os.Chdir(path) != nil{
					fmt.Println("cd: " + path + ": No such file or directory")
				}
			}
			
		} else {
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
