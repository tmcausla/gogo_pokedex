package main
import "fmt"

func commandInspect(c *config, keyword string) error {

	// TO DO:
	// Build out pokemon struct to include stats and types

	pokemon := c.pokedex[keyword]
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	return nil
}