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

func checkNeighbor(lines *[]string, seen *map[string]int, selfCol, selfRow, neighborCol, neighborRow int) {
	if (*lines)[neighborRow][neighborCol] == '@' {
		self := fmt.Sprintf("%d-%d", selfCol, selfRow)
		neighbor := fmt.Sprintf("%d-%d", neighborCol, neighborRow)
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
	for row, line := range lines {
		for col, r := range line {
			if r != '@' {
				continue
			}

			loc := fmt.Sprintf("%d-%d", col, row)
			_, found := seen[loc]
			if !found {
				seen[loc] = 0
			}

			// check right
			if col+1 < rowLength {
				checkNeighbor(&lines, &seen, col, row, col+1, row)
			}

			// check down
			if row+1 < linesLength {
				checkNeighbor(&lines, &seen, col, row, col, row+1)

				// check down left
				if col > 0 {
					checkNeighbor(&lines, &seen, col, row, col-1, row+1)
				}

				// check down right
				if col+1 < rowLength {
					checkNeighbor(&lines, &seen, col, row, col+1, row+1)
				}
			}

			if neighbors := seen[loc]; neighbors < 4 {
				pass++
			}
		}
	}

	return pass, nil
}

func day4part2(lines []string) (int, error) {
	pass := 0

	seen := make(map[string]int)
	qVisited := make(map[string]bool)
	queue := Queue[Loc]{}
	linesLength := len(lines)
	rowLength := len(lines[0])

	for row, line := range lines {
		for col, r := range line {
			if r != '@' {
				continue
			}

			loc := fmt.Sprintf("%d-%d", col, row)
			_, found := seen[loc]
			if !found {
				seen[loc] = 0
			}

			// check right
			if col+1 < rowLength {
				checkNeighbor(&lines, &seen, col, row, col+1, row)
			}

			// check down
			if row+1 < linesLength {
				checkNeighbor(&lines, &seen, col, row, col, row+1)

				// check down left
				if col > 0 {
					checkNeighbor(&lines, &seen, col, row, col-1, row+1)
				}

				// check down right
				if col+1 < rowLength {
					checkNeighbor(&lines, &seen, col, row, col+1, row+1)
				}
			}

			if neighbors := seen[loc]; neighbors < 4 {
				queue.enqueue(Loc{col, row})
				qVisited[loc] = true
			}

		}
	}

	queueNeighbors := func(col, row int) {
		dirs := []Loc{
			{-1, -1}, // up left
			{0, -1},  // up
			{1, -1},  // up right
			{-1, 0},  // left
			{1, 0},   // right
			{-1, 1},  // down left
			{0, 1},   // down
			{1, 1},   // down right
		}

		for _, dir := range dirs {
			checkCol := col + dir.col
			checkRow := row + dir.row
			if checkCol >= 0 && checkCol < linesLength && checkRow >= 0 && checkRow < rowLength {
				loc := fmt.Sprintf("%d-%d", checkCol, checkRow)
				if _, found := seen[loc]; found {
					seen[loc]--
					if seen[loc] < 4 && !qVisited[loc] {
						qVisited[loc] = true
						queue.enqueue(Loc{checkCol, checkRow})
					}
				}
			}
		}
	}

	for !queue.isEmpty() {
		pass++
		loc := queue.dequeue()
		delete(seen, fmt.Sprintf("%d-%d", loc.col, loc.row))
		queueNeighbors(loc.col, loc.row)
	}

	return pass, nil
}
