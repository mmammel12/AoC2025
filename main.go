package main

import (
	"fmt"
	"os"
	"strings"
)

func getCommands() map[string]func([]string) (int, error) {
	return map[string]func([]string) (int, error){
		"1-1": day1part1,
		"1-2": day1part2,
	}
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Error: not enough arguments")
		fmt.Println("First arg is day [1,2,3,...,25]")
		fmt.Println("Second arg is part [1, 2]")
		fmt.Println("Optional third arg to use test file (test.txt) [t]")
		fmt.Println("Optional fourth arg is to specify the test file name")
		fmt.Println("\texample: './AoC2025 1 1 t hello' will run day 1 part 1 against hello.txt")
		os.Exit(1)
	}

	fileName := "input.txt"
	if len(args) == 4 {
		fileName = "test.txt"
	} else if len(args) == 5 {
		fileName = fmt.Sprintf("%v.txt", args[4])
	}
	path := fmt.Sprintf("./day%v/%v", args[1], fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	commands := getCommands()
	commandName := fmt.Sprintf("%v-%v", args[1], args[2])

	if fn, exists := commands[commandName]; exists {
		ans, err := fn(lines[:len(lines)-1])
		if err != nil {
			fmt.Printf("Error in part 1: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Answer: %v\n", ans)
	}

	os.Exit(0)
}
