package main

import (
	"strconv"
	"strings"
)

func day6(lines []string, part int) (int, error) {
	if part == 1 {
		return day6part1(lines)
	}

	return day6part2(lines)
}

func day6part1(lines []string) (int, error) {
	pass := 0

	n := len(lines) - 1
	operands := strings.Fields(lines[n])
	answers := make([]int, len(operands))

	for i := range answers {
		if operands[i] == "+" {
			answers[i] = 0
		} else {
			answers[i] = 1
		}
	}

	for row := range lines {
		if row == n {
			break
		}

		numStrs := strings.Fields(lines[row])
		for col, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			if operands[col] == "+" {
				answers[col] += num
			} else {
				answers[col] *= num
			}
		}
	}

	for _, ans := range answers {
		pass += ans
	}

	return pass, nil
}

func day6part2(lines []string) (int, error) {
	pass := 0
	n := len(lines) - 1
	operands := strings.Fields(lines[n])
	answers := make([]int, len(operands))
	for i := range answers {
		if operands[i] == "+" {
			answers[i] = 0
		} else {
			answers[i] = 1
		}
	}

	answerCol := len(answers) - 1
	skip := false
	for col := len(lines[0]) - 1; col >= 0; col-- {
		if skip {
			skip = false
			answerCol--
			continue
		}

		valSeen := false

		operand := operands[answerCol]
		numStr := ""
		for row := range lines {
			if row == n {
				if lines[row][col] == '+' || lines[row][col] == '*' {
					skip = true
				}
				break
			}

			val := string(lines[row][col])
			if val == " " {
				if !valSeen {
					numStr += "0"
				}
			} else {
				valSeen = true
				numStr += val
			}
		}

		num, _ := strconv.Atoi(numStr)
		if operand == "+" {
			answers[answerCol] += num
		} else {
			answers[answerCol] *= num
		}
	}

	for _, ans := range answers {
		pass += ans
	}

	return pass, nil
}
