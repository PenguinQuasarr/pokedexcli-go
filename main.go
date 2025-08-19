package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func repl() {

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		input.Scan()
		command := cleanInput(input.Text())
		fmt.Printf("Your command was: %s\n", command[0])
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	trimedAndLower := strings.ToLower(strings.TrimSpace(text))
	words := strings.Split(trimedAndLower, " ")
	return words
}

func main() {

	repl()

}
