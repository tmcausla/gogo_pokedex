package main
import (
    "github.com/tmcausla/gogo_pokedex/internal/pokeapi"
)

func main() {
    base_url := "https://pokeapi.co/api/v2/location-area/"
    replConfig := &config{
        cmdList: map[string]cliCommand{
            "help": {
                name: 	 	 "help",
                description: "Displays a help message",
                callback: 	 commandHelp,
            },
            "exit": {
                name: 		 "exit",
                description: "Exit the Pokedex",
                callback: 	 commandExit,
            },
            "map": {
                name:        "map",
                description: "Displays the next 20 area locations",
                callback:    commandMap,
            },
            "mapb": {
                name:        "mapb",
                description: "Displays the previous 20 area locations",
                callback:    commandMapB,
            },
        },
        nextLocationsURL: &base_url,
        prevLocationsURL: nil,
    }
    replConfig.client = pokeapi.NewClient()

    
    startRepl(replConfig)
}
