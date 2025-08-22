package main

import (
	"fmt"
	"os"
	"github.com/PenguinQuasarr/pokedexcli-go/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(map[string]cliCommand, *Config) error
}

type Config struct {
	next string
	previous string
}

func commandMap(commands map[string]cliCommand, con *Config) error {

	fmt.Println(con.next)
	locations, err := pokeapi.GetPokeData(con.next)
	if err != nil {
		fmt.Println(err)
	}

	con.next = locations.Next
	con.previous = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return err
}

func commandMapb(commands map[string]cliCommand, con *Config) error {

	fmt.Println(con.previous)
	locations, err := pokeapi.GetPokeData(con.previous)
	if locations.Previous == "" {
		fmt.Println("On first page")
		return nil
	}
	if err != nil {
		fmt.Println(err)
	}

	con.next = locations.Next
	con.previous = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return err
}

func commandExit(commands map[string]cliCommand, con *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand, con *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}


func GetCommands() map[string]cliCommand {

	commands := map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Show location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show previous areas",
			callback: commandMapb,
		},
	}

	return commands
}




