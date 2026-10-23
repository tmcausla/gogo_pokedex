package main
import "fmt"

func commandInspect(c *config, keyword string) error {
	pokemon, ok := c.pokedex[keyword]
	if !ok {
		fmt.Printf("you haven't caught that pokemon\nuse 'catch %s' to throw a Pokeball at it!\n", keyword)
		return nil
	}
	
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, pokestat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokestat.Stat.Name, pokestat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range pokemon.TypeList {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}

	return nil
}