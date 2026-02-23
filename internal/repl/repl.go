package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NicholasGSwan/pokedexcli/internal/models"
	"github.com/NicholasGSwan/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand

type config struct {
	next string
	prev string
	arg  string
}

var cache *pokecache.Cache

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists areas of map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of map areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all pokemon in Location Area. Takes location area name as an argument.",
			callback:    commandPokemonInLocArea,
		},
	}
	cache = pokecache.NewCache(time.Second * 5)
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	strArr := strings.Fields(text)
	return strArr
}

func StartRepl() {
	c := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		command, ok := commands[words[0]]
		if len(words) > 1 {
			c.arg = words[1]
		}
		if ok {
			command.callback(c)
		} else {
			fmt.Println("Unknown command")
		}
		// fmt.Printf("Your command was: %s\n", words[0])

	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

// Map command; gets next map on successive calls
func commandMap(c *config) error {
	var url string
	if c.next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = c.next
	}

	locArea, err := getMapResults(url)
	if err != nil {
		return err
	}
	printLocationAreas(locArea, c)
	return nil
}

// Map back command
func commandMapb(c *config) error {
	var url string
	if c.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = c.prev
	}

	locArea, err := getMapResults(url)
	if err != nil {
		return err
	}
	printLocationAreas(locArea, c)
	return nil
}

func commandPokemonInLocArea(c *config) error {
	var url string
	if c.arg == "" {
		fmt.Println("Command requires an argument")
	} else {
		url = "https://pokeapi.co/api/v2/location-area/" + c.arg
	}
	LA, err := getLocationDetails(url)
	if err != nil {
		return err
	}
	for _, pokeEnc := range LA.PokemonEncounters {
		fmt.Println(pokeEnc.Pokemon.Name)
	}
	return nil
}

func getMapResults(url string) (models.LocationAreaGetResult, error) {
	var locArea models.LocationAreaGetResult
	data, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &locArea)
		if err != nil {
			return models.LocationAreaGetResult{}, err
		}
		return locArea, err
	}
	res, err := http.Get(url)
	if err != nil {
		return models.LocationAreaGetResult{}, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locArea); err != nil {
		return models.LocationAreaGetResult{}, err
	}
	return locArea, nil
}

func printLocationAreas(locArea models.LocationAreaGetResult, c *config) {
	c.next = locArea.Next
	c.prev = locArea.Previous
	areas := locArea.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}
}

func getLocationDetails(url string) (models.LocationArea, error) {
	var LA models.LocationArea
	data, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &LA)
		if err != nil {
			return models.LocationArea{}, err
		}
		return LA, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return models.LocationArea{}, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&LA); err != nil {
		return models.LocationArea{}, err
	}
	return LA, nil
}
