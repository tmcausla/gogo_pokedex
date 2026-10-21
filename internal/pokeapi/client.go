package pokeapi

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}

func (c Client) ListLocations(url string) (LocationList, error) {
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

	return locations, nil
}