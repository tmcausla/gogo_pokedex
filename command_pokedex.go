package main
import "fmt"

func commandPokedex(c *config, keyword string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("you haven't caught any Pokemon yet\nget out there and toss some balls yo")
		return nil
	}

	fmt.Println("Your pokedex:")
	for pokemon := range c.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}