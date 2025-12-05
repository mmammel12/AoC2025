package main

import (
	"fmt"
)

func day4(lines []string, part int) (int, error) {
	if part == 1 {
		return day4part1(lines)
	}

	return day4part2(lines)
}

func checkNeighbor(lines *[]string, seen *map[string]int, selfX, selfY, neighborX, neighborY int) {
	if (*lines)[neighborY][neighborX] == '@' {
		self := fmt.Sprintf("%d-%d", selfX, selfY)
		neighbor := fmt.Sprintf("%d-%d", neighborX, neighborY)
		_, rFound := (*seen)[neighbor]
		if !rFound {
			(*seen)[neighbor] = 0
		}
		(*seen)[self]++
		(*seen)[neighbor]++
	}
}

func day4part1(lines []string) (int, error) {
	pass := 0

	seen := make(map[string]int)
	linesLength := len(lines)
	rowLength := len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			if r != '@' {
				continue
			}

			loc := fmt.Sprintf("%d-%d", x, y)
			_, found := seen[loc]
			if !found {
				seen[loc] = 0
			}

			// check right
			if x+1 < rowLength {
				checkNeighbor(&lines, &seen, x, y, x+1, y)
			}

			// check down
			if y+1 < linesLength {
				checkNeighbor(&lines, &seen, x, y, x, y+1)

				// check down left
				if x > 0 {
					checkNeighbor(&lines, &seen, x, y, x-1, y+1)
				}

				// check down right
				if x+1 < rowLength {
					checkNeighbor(&lines, &seen, x, y, x+1, y+1)
				}
			}

			if neighbors := seen[loc]; neighbors < 4 {
				pass++
			}
		}
	}

	return pass, nil
}

type Loc struct {
	x int
	y int
}
type Queue []Loc

func (q *Queue) enqueue(x, y int) {
	*q = append(*q, Loc{x: x, y: y})
}

func (q *Queue) dequeue() Loc {
	loc := (*q)[0]
	*q = (*q)[1:]
	return loc
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func day4part2(lines []string) (int, error) {
	pass := 0

	seen := make(map[string]int)
	qVisited := make(map[string]bool)
	queue := Queue{}
	linesLength := len(lines)
	rowLength := len(lines[0])

	for y, line := range lines {
		for x, r := range line {
			if r != '@' {
				continue
			}

			loc := fmt.Sprintf("%d-%d", x, y)
			_, found := seen[loc]
			if !found {
				seen[loc] = 0
			}

			// check right
			if x+1 < rowLength {
				checkNeighbor(&lines, &seen, x, y, x+1, y)
			}

			// check down
			if y+1 < linesLength {
				checkNeighbor(&lines, &seen, x, y, x, y+1)

				// check down left
				if x > 0 {
					checkNeighbor(&lines, &seen, x, y, x-1, y+1)
				}

				// check down right
				if x+1 < rowLength {
					checkNeighbor(&lines, &seen, x, y, x+1, y+1)
				}
			}

			if neighbors := seen[loc]; neighbors < 4 {
				queue.enqueue(x, y)
				qVisited[loc] = true
			}

		}
	}

	queueNeighbors := func(x, y int) {
		dirs := []Loc{
			{x: -1, y: -1}, // up left
			{x: 0, y: -1},  // up
			{x: 1, y: -1},  // up right
			{x: -1, y: 0},  // left
			{x: 1, y: 0},   // right
			{x: -1, y: 1},  // down left
			{x: 0, y: 1},   // down
			{x: 1, y: 1},   // down right
		}

		for _, dir := range dirs {
			checkX := x + dir.x
			checkY := y + dir.y
			if checkX >= 0 && checkX < linesLength && checkY >= 0 && checkY < rowLength {
				loc := fmt.Sprintf("%d-%d", checkX, checkY)
				if _, found := seen[loc]; found {
					seen[loc]--
					if seen[loc] < 4 && !qVisited[loc] {
						qVisited[loc] = true
						queue.enqueue(checkX, checkY)
					}
				}
			}
		}
	}

	for !queue.isEmpty() {
		pass++
		loc := queue.dequeue()
		delete(seen, fmt.Sprintf("%d-%d", loc.x, loc.y))
		queueNeighbors(loc.x, loc.y)
	}

	return pass, nil
}
