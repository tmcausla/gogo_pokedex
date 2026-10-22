package main
import "fmt"

func commandMapB(c *config, keyword string) error {
	if c.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationList, err := c.client.GetLocationList(*c.prevLocationsURL)
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