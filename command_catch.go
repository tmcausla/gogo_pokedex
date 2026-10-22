package main
import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, keyword string) error {
	if keyword == "" {
		fmt.Println("no pokemon provided\nuse 'explore {location}' to see a list of available pokemon to catch")
		return nil
	}

	if _, ok := c.pokedex[keyword]; ok {
		fmt.Printf("you already caught %s!\nuse 'inspect %s' to check out its stats\n", keyword, keyword)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", keyword)

	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	pokemon, err := c.client.GetPokemon(baseURL + keyword)
	if err != nil {
		fmt.Print("something went wrong. double check spelling and try again\n")
		return nil
	}
	
	catchChance := 50 - pokemon.BaseExp / 4
	if (catchChance < 5) {
		catchChance = 5
	}
	isCaught := rand.Intn(100) < catchChance

	if isCaught {
		fmt.Printf("%s was caught!\n", keyword)
		c.pokedex[keyword] = pokemon
		return nil
	} else {
		fmt.Printf("%s escaped!\n", keyword)
		return nil
	}
}