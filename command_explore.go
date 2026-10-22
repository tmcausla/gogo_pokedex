package main
import "fmt"

func commandExplore(c *config, keyword string) error {
	if keyword == "" {
		fmt.Println("no location provided\nuse 'map' and 'mapb' to see a list of available locations to explore")
		return nil
	}

	baseURL := "https://pokeapi.co/api/v2/location-area/"
	encounterList, err := c.client.GetEncounterList(baseURL + keyword)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon: \n", keyword)
	for _, encounter := range encounterList.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}