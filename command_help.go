package main
import "fmt"

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range c.cmdList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}