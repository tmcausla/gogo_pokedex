package pokeapi
import (
	"fmt"
	"io"
	"encoding/json"
)

type Pokemon struct {
	Name 	string `json:"name"`
	BaseExp int    `json:"base_experience`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
}

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	if entry, ok := c.cache.Get(url); ok {
		fmt.Print("cache hit! getting stored data...\n\n")
		var entryPokemon Pokemon
		if err := json.Unmarshal(entry, &entryPokemon); err != nil {
			return Pokemon{}, fmt.Errorf("error loading Pokemon from cache: %w", err)
		}
		return entryPokemon, nil
	}

	fmt.Print("making http request for data...\n\n")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("error loading json data into Pokemon struct: %w", err)
	}

	c.cache.Add(url, data)
	return pokemon, nil
}