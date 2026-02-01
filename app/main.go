package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
			fmt.Fprintln(os.Stderr, "Error reading inputß: ", err)
			os.Exit(1)
		}

		//Clean saces at the beginning and end ex. echo hello world
		cleanInput := strings.TrimSpace(input)
		parts, err := parseInput(cleanInput)

		if err != nil {
			fmt.Println("Error parsing input:", err)
			continue
		}
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


func parseInput(input string)([]string,error){
	var args []string
	var currentArg strings.Builder 
	inSingleQuote := false
	inDoubleQuote := false
	escape := false

	for _, r := range input{
		if escape {
			if inDoubleQuote{
				switch r{
					case '$', '`', '"', '\\', '\n':
                    //If one of these specials, we kill the backslash and just write the character
                    currentArg.WriteRune(r)
                default:
                    //We write the backslash again and the character if the character that follows the backlash is not special
                    currentArg.WriteRune('\\') 
                    currentArg.WriteRune(r)
				}
			}else{
				//outside double quotes, a backslash ALWAYS escapes the next character
                currentArg.WriteRune(r)
			}

			escape = false
			continue
		}

		//if there is a backslash
		if r == '\\'{ //in other words we are just checking for a single \
			if inSingleQuote{
				currentArg.WriteRune(r)
			}else{
				escape = true
			}
			continue
		}
		
		//if there is a single quote
		if r == '\''{ //in other words we are just checking for ' 
			if inDoubleQuote{
				currentArg.WriteRune(r)
			}else{
				inSingleQuote = !inSingleQuote
			}
			continue
		}

		if r == '"'{
			if inSingleQuote{
				currentArg.WriteRune(r)
			}else{
				inDoubleQuote = !inDoubleQuote
			}
			continue
		}

		if r == ' '{
			if !inSingleQuote && !inDoubleQuote{
				if currentArg.Len() > 0{
					args = append(args, currentArg.String())
					currentArg.Reset()
				}
				continue
			}
			//if we are in quotes then we do need the space
			currentArg.WriteRune(r)
			continue
		}
		currentArg.WriteRune(r)
	}

	if currentArg.Len() > 0{
		args = append(args, currentArg.String())
	}

	if inDoubleQuote || inSingleQuote{
		return nil, fmt.Errorf("unclosed quote")
	}

	return args, nil

}