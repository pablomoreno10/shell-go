package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
)

var _ = fmt.Print

func main() {

	exit := "exit"
	echo := "echo"
	builtin := "type"

	for{
		fmt.Print("$ ")

		//read user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		//invalid input
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input: " ,err)
			os.Exit(1)
		}

		//commands
		if  strings.Fields(strings.TrimSpace(command))[0] == builtin{
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			arg := strings.Join(slice, " ")
			if (arg == exit) || (arg == builtin) || (arg == echo){
				fmt.Println(arg + " is a shell builtin")
			}else{
				slice := strings.Fields(strings.TrimSpace(command))[1:]
				arg := strings.Join(slice, " ")
				path, err := exec.LookPath(arg)
				if err != nil {
					fmt.Println(strings.TrimSpace(arg)+ ": not found")
				}else{
					fmt.Println(arg + " is " +  path)
				}
			}
		}else if strings.Fields(strings.TrimSpace(command))[0] == echo{
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			arg := strings.Join(slice, " ")
			fmt.Println(arg)
		}else if strings.TrimSpace(command) == exit{
			os.Exit(0)
		}else{
			executable := strings.Fields(strings.TrimSpace(command))[0]
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			_ , err := exec.LookPath(executable)
			if err != nil {
				fmt.Println(strings.TrimSpace(command) + ": command not found")
			}else{
				cmd := exec.Command(executable, slice...)	//ellipsis to pass all arguments of the slice into the variadic parameter
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("Error running command: %v\n", err)
				}else{
					fmt.Printf(string(output))
				}
			}
		}
	}
}
