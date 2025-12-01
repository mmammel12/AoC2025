package main

import (
	"math"
	"strconv"
)

func day1part1(lines []string) (int, error) {
	current := 50
	pass := 0

	for _, ins := range lines {
		direction := string(ins[0])
		distance, err := strconv.Atoi(ins[1:])
		if err != nil {
			return 0, err
		}

		if direction == "R" {
			current = (current + distance) % 100
		} else if distance > current {
			current = (current - distance) + 100
		} else {
			current -= distance
		}

		if current == 0 {
			pass++
		}
	}

	return pass, nil
}

func day1part2(lines []string) (int, error) {
	current := 50
	pass := 0

	for _, ins := range lines {
		direction := string(ins[0])
		distance, err := strconv.Atoi(ins[1:])
		if err != nil {
			return 0, err
		}

		rotations := int(math.Abs(math.Floor(float64(distance) / 100.0)))
		distance %= 100
		pass += rotations

		if direction == "R" {
			if (current + distance) > 100 {
				pass++
			}
			current = (current + distance) % 100
		} else {
			if distance > current {
				if current != 0 {
					pass++
				}
				current = (current - distance) + 100
			} else {
				current -= distance
			}
		}

		if current == 0 {
			pass++
		}
	}

	return pass, nil
}
