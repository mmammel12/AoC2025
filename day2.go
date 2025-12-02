package main

import (
	"strconv"
	"strings"
)

func day2part1(lines []string) (int, error) {
	pass := 0

	for id := range strings.SplitSeq(lines[0], ",") {
		splitID := strings.Split(id, "-")

		startStr := splitID[0]
		current := startStr[:len(startStr)/2]

		startID, _ := strconv.Atoi(startStr)
		endID, _ := strconv.Atoi(splitID[1])

		for {
			currentNum, _ := strconv.Atoi(current + current)
			if currentNum >= startID && currentNum <= endID {
				pass += currentNum
			}
			if currentNum > endID {
				break
			}
			temp, _ := strconv.Atoi(current)
			current = strconv.Itoa(temp + 1)
		}
	}

	return pass, nil
}

func day2part2(lines []string) (int, error) {
	powers := []int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
		100000000000,
	}
	mods := [][]int{
		{},
		{},
		{11},
		{111},
		{101, 111},
		{11111},
		{1001, 111111, 10101},
		{1111111},
		{10001, 11111111, 1010101},
		{111111111, 1001001},
		{100001, 1111111111, 101010101},
	}
	pass := 0

	for id := range strings.SplitSeq(lines[0], ",") {
		splitID := strings.Split(id, "-")

		startStr := splitID[0]
		endStr := splitID[1]

		startID, _ := strconv.Atoi(startStr)
		endID, _ := strconv.Atoi(endStr)

		for length := len(startStr); length <= len(endStr); length++ {
			seen := make(map[int]struct{})
			start := max(startID, powers[length-1])
			end := min(endID, powers[length])
			for _, mod := range mods[length] {
				current := start - (start % mod)
				for current <= end {
					_, found := seen[current]
					if current >= start && !found {
						pass += current
						seen[current] = struct{}{}
					}
					current += mod
				}
			}
		}
	}

	return pass, nil
}
