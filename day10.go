package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day10(lines []string, part int) (int, error) {
	if part == 1 {
		return day10part1(lines)
	}

	return day10part2(lines)
}

type machine struct {
	diagram string
	buttons [][]int
	joltage []int
}

func createMachines(lines []string) []machine {
	n := len(lines)
	machines := make([]machine, n)
	r := regexp.MustCompile(`\[([.#]+)\]|\(([0-9,]+)\)|\{([0-9,]+)\}`)
	for i, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)

		diagram := matches[0][1]

		buttons := make([][]int, len(matches)-2)
		for j := 1; j < len(matches)-1; j++ {
			parts := strings.Split(matches[j][2], ",")
			b := make([]int, len(parts))
			for k, part := range parts {
				num, _ := strconv.Atoi(part)
				b[k] = num
			}
			buttons[j-1] = b
		}

		joltageStr := matches[len(matches)-1][3]
		parts := strings.Split(joltageStr, ",")
		joltage := make([]int, len(parts))
		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			joltage[j] = num
		}

		machines[i] = machine{diagram, buttons, joltage}
	}

	return machines
}

func applyDiagramButtons(str string, buttons []int) string {
	runes := []rune(str)
	for _, b := range buttons {
		if runes[b] == '.' {
			runes[b] = '#'
		} else {
			runes[b] = '.'
		}
	}

	return string(runes)
}

func day10part1(lines []string) (int, error) {
	presses := 0

	machines := createMachines(lines)
	for _, m := range machines {
		q := Queue[string]{}
		visited := make(map[string]bool)

		startStr := strings.Repeat(".", len(m.diagram))
		dp := make(map[string]int)
		dp[startStr] = 0
		q = append(q, startStr)

		targetFound := false
		for !q.isEmpty() && !targetFound {
			current := q.dequeue()
			visited[current] = true
			for _, btns := range m.buttons {
				inc := dp[current] + 1
				newStr := applyDiagramButtons(current, btns)

				if newStr == m.diagram {
					presses += inc
					targetFound = true
					break
				}

				if _, exists := dp[newStr]; !exists {
					dp[newStr] = inc
				}

				if _, exists := visited[newStr]; !exists {
					q.enqueue(newStr)
				}
			}
		}
	}

	return presses, nil
}

func day10part2(lines []string) (int, error) {
	presses := 0

	return presses, nil
}
