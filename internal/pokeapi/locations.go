package pokeapi
import (
	"fmt"
	"io"
	"encoding/json"
)

type LocationList struct {
	Count 	 int 		`json:"count"`
	Next 	 *string 	`json:"next"`
	Previous *string 	`json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) GetLocationList(url string) (LocationList, error) {
	if entry, ok := c.cache.Get(url); ok {
		fmt.Print("cache hit! getting stored data...\n\n")
		var entryLocations LocationList
		if err := json.Unmarshal(entry, &entryLocations); err != nil {
			return LocationList{}, fmt.Errorf("error loading LocationList from cache: %w", err)
		}
		return entryLocations, nil
	}

	fmt.Print("making http request for data...\n\n")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationList{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{}, err
	}
	if res.StatusCode > 299 {
		return LocationList{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var locations LocationList
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationList{}, fmt.Errorf("error loading json data into LocationList struct: %w", err)
	}

	c.cache.Add(url, data)
	return locations, nil
}