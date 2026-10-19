package main
import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

var commandList map[string]cliCommand
func init() {
	commandList = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
	}
}

func cleanInput(text string) []string {
    cleanText := strings.ToLower(text)
	return strings.Fields(cleanText)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        if !scanner.Scan() {
            break
        }

        words := cleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }

        firstWord := words[0]
        if cmd, ok := commandList[firstWord]; ok {
            err := cmd.callback()
            if err != nil {
                fmt.Println("Unknown command")
            }
        }
    }
}



