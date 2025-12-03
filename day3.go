package main

import (
	"math"
	"strconv"
)

func day3part1(lines []string) (int, error) {
	pass := 0

	for _, bank := range lines {
		length := len(bank)
		tens := '0'
		ones := '0'

		for i, b := range bank {
			if rune(b) > tens && i < length-1 {
				tens = rune(b)
				ones = '0'
			} else if rune(b) > ones {
				ones = rune(b)
			}
		}

		num, err := strconv.Atoi(string(tens) + string(ones))
		if err != nil {
			return 0, err
		}
		pass += num
	}

	return pass, nil
}

func rtoi(r rune) int {
	return int(r - '0')
}

func day3part2(lines []string) (int, error) {
	pass := 0

	for _, bank := range lines {
		n := len(bank)
		left := 0
		right := 0
		remaining := 12
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
