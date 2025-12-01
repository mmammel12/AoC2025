package main

func getCommands() map[string]func([]string) (int, error) {
	return map[string]func([]string) (int, error){
		"1-1": day1part1,
		"1-2": day1part2,
	}
}
