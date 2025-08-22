package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	trimedAndLower := strings.ToLower(strings.TrimSpace(text))
	words := strings.Split(trimedAndLower, " ")
	return words
}


func repl() {

	input := bufio.NewScanner(os.Stdin)
	commands := GetCommands()
	config := Config{}
	config.next = "https://pokeapi.co/api/v2/location-area/"

	for {
		fmt.Print("pokedex > ")
		input.Scan()
		command := cleanInput(input.Text())

		cmd, ok := commands[command[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			cmd.callback(commands, &config)
		}
	}
}

func main() {
	repl()
}
