package main
import (
    "github.com/tmcausla/gogo_pokedex/internal/pokeapi"
    "time"
)

func main() {
    base_url := "https://pokeapi.co/api/v2/location-area/"
    cacheInterval := 10 * time.Second
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
            "explore": {
                name:        "explore {location}",
                description: "Displays a list of pokemon found at a location",
                callback:    commandExplore,
            },
            "catch": {
                name:        "catch {pokemon}",
                description: "Attempt to catch a Pokemon",
                callback:    commandCatch,
            },
            "inspect": {
                name:        "inspect {pokemon}",
                description: "Displays stats of a caught Pokemon",
                callback:    commandInspect,
            },
            "pokedex": {
                name:        "pokedex",
                description: "Displays a list of all pokemon you've caught",
                callback:    commandPokedex,
            },
        },
        nextLocationsURL: &base_url,
        prevLocationsURL: nil,
        pokedex: map[string]pokeapi.Pokemon{},
    }
    replConfig.client = pokeapi.NewClient(cacheInterval) 

    
    startRepl(replConfig)
}
