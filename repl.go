package main
import (
	"strings"
	"fmt"
	"os"
	"bufio"
    "github.com/tmcausla/gogo_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func(*config, string) error
}

type config struct {
	cmdList          map[string]cliCommand
    nextLocationsURL *string
    prevLocationsURL *string
    client           pokeapi.Client
    pokedex          map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
    cleanText := strings.ToLower(text)
	return strings.Fields(cleanText)
}

func startRepl(c *config) {
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
        keyword := ""
        if len(words) > 1 {
            keyword = words[1]
        }

        if cmd, ok := c.cmdList[firstWord]; ok {
            err := cmd.callback(c, keyword)
            if err != nil {
                fmt.Println("Unknown command")
            }
        } else {
            fmt.Println("Unknown command\nType 'help' to see list of valid commands")
        }
    }
}



