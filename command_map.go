package main
import "fmt"

func commandMap(c *config) error {
	if c.nextLocationsURL == nil {
		fmt.Println("you're on the last page")
		return nil
	}

	locationList, err := c.client.ListLocations(*c.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	c.nextLocationsURL = locationList.Next
	c.prevLocationsURL = locationList.Previous

	return nil
}