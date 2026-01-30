package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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

		//type builtin
		if  strings.Fields(strings.TrimSpace(command))[0] == builtin{
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			output := strings.Join(slice, " ")
			if (output == exit) || (output == builtin) || (output == echo){
				fmt.Println(output + " is a shell builtin")
			}else{
				slice := strings.Fields(strings.TrimSpace(command))[1:]
				output := strings.Join(slice, " ")
				fmt.Println(strings.TrimSpace(output)+ ": not found")
			}
		}else if strings.Fields(strings.TrimSpace(command))[0] == echo{
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			output := strings.Join(slice, " ")
			fmt.Println(output)
		}else if strings.TrimSpace(command) == exit{
			os.Exit(0)
		}else{
			fmt.Println(strings.TrimSpace(command) + ": command not found")
		}
	}
}
