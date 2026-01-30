package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var _ = fmt.Print

func main() {
	var exit,echo string
	exit = "exit"
	echo = "echo"
	for{
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input: " ,err)
			os.Exit(1)
		}

		if strings.TrimSpace(command) == exit{
			os.Exit(0)
		}

		if  strings.Fields(strings.TrimSpace(command))[0] == echo{
			slice := strings.Fields(strings.TrimSpace(command))[1:]
			output := strings.Join(slice, " ")
			fmt.Println(output)
		}else{
			fmt.Println(strings.TrimSpace(command)+ ": command not found")
		}
	}
	

}
