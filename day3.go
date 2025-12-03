package main

import (
	"math"
)

func day3(lines []string, part int) (int, error) {
	if part == 1 {
		return day3part1(lines, 2)
	}

	return day3part1(lines, 12)
}

func day3part1(lines []string, length int) (int, error) {
	pass := 0

	for _, bank := range lines {
		n := len(bank)
		left := 0
		right := 0
		remaining := length
		temp := 0

		for remaining > 0 {
			current := '0'
			for right <= n-remaining {
				if current < rune(bank[right]) {
					current = rune(bank[right])
					left = right
				}
				right++
			}
			remaining--
			left++
			right = left
			temp += rtoi(current) * int(math.Pow10(remaining))
		}
		pass += temp
	}

	return pass, nil
}

func rtoi(r rune) int {
	return int(r - '0')
}
