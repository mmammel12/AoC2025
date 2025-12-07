package main

import (
	"slices"
	"strings"
)

func day7(lines []string, part int) (int, error) {
	if part == 1 {
		return day7part1(lines)
	}

	return day7part2(lines)
}

func day7part1(lines []string) (int, error) {
	pass := 0

	startCol := strings.IndexRune(lines[0], 'S')
	splits := make(map[int]bool)
	splits[startCol] = true

	for i, line := range lines {
		// splitters only on even lines
		// start is on line 0
		if i%2 != 0 || i == 0 {
			continue
		}

		for col := range splits {
			if line[col] == '^' {
				pass++
				delete(splits, col)
				// splitters are never on the edge
				splits[col-1] = true
				splits[col+1] = true
			}
		}
	}

	return pass, nil
}

func day7part2(lines []string) (int, error) {
	pass := 0

	n := len(lines)
	startCol := strings.IndexRune(lines[0], 'S')
	root := Loc{startCol, 2}
	ends := make(map[int]int)
	seen := make(map[Loc]int)
	seen[root] = 1
	q := Queue[Loc]{root}

	for !q.isEmpty() {
		current := q.dequeue()
		paths := seen[current]

		leftHit := false
		rightHit := false
		col := current.col
		for row := current.row + 2; row < n; row += 2 {
			if !leftHit && lines[row][col-1] == '^' {
				leftHit = true
				loc := Loc{col - 1, row}
				if _, exists := seen[loc]; !exists {
					seen[loc] = 0
					q.enqueue(loc)
				}
				seen[loc] += paths
			}
			if !rightHit && lines[row][col+1] == '^' {
				rightHit = true
				loc := Loc{col + 1, row}

				if _, exists := seen[loc]; !exists {
					seen[loc] = 0
					q.enqueue(loc)
				}
				seen[loc] += paths
			}

			if leftHit && rightHit {
				break
			}
		}

		if !leftHit {
			if _, exists := ends[col-1]; !exists {
				ends[col-1] = 0
			}
			ends[col-1] += paths
		}
		if !rightHit {
			if _, exists := ends[col+1]; !exists {
				ends[col+1] = 0
			}
			ends[col+1] += paths
		}

		// sort queue to make sure lower rows are not processed too soon
		slices.SortFunc(q, func(a, b Loc) int {
			return a.row - b.row
		})
	}

	for _, paths := range ends {
		pass += paths
	}

	return pass, nil
}
