package pokeapi
import (
	"fmt"
	"io"
	"encoding/json"
)

type EncounterList struct {
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Encounter `json:"pokemon"`
}

type Encounter struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) GetEncounterList(url string) (EncounterList, error) {
	if entry, ok := c.cache.Get(url); ok {
		fmt.Print("cache hit! getting stored data...\n\n")
		var entryEncounters EncounterList
		if err := json.Unmarshal(entry, &entryEncounters); err != nil {
			return EncounterList{}, fmt.Errorf("error loading EncounterList from cache: %w", err)
		}
		return entryEncounters, nil
	}

	fmt.Print("making http request for data...\n\n")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return EncounterList{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EncounterList{}, err
	}
	if res.StatusCode > 299 {
		return EncounterList{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var encounters EncounterList
	if err := json.Unmarshal(data, &encounters); err != nil {
		return EncounterList{}, fmt.Errorf("error loading json data into EncounterList struct: %w", err)
	}

	c.cache.Add(url, data)
	return encounters, nil
}