package main

import (
	"math"
	"strconv"
	"strings"
)

func day9(lines []string, part int) (int, error) {
	if part == 1 {
		return day9part1(lines)
	}

	return day9part2(lines)
}

func getLoc(line string) Loc {
	parts := strings.Split(line, ",")
	col, _ := strconv.Atoi(parts[0])
	row, _ := strconv.Atoi(parts[1])

	return Loc{col, row}
}

func day9part1(lines []string) (int, error) {
	maxArea := 0

	n := len(lines)
	left := 0
	for left < n-1 {
		leftLoc := getLoc(lines[left])
		right := left + 1
		for right < n {
			rightLoc := getLoc(lines[right])

			width := max(max(leftLoc.col, rightLoc.col)-min(leftLoc.col, rightLoc.col), 1) + 1
			height := max(max(leftLoc.row, rightLoc.row)-min(leftLoc.row, rightLoc.row), 1) + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}

			right++
		}
		left++
	}

	return maxArea, nil
}

type rectangle struct {
	left   int
	right  int
	top    int
	bottom int
}

func day9part2(lines []string) (int, error) {
	maxArea := 0
	n := len(lines)
	vertices := []Loc{}
	for _, line := range lines {
		loc := getLoc(line)
		vertices = append(vertices, loc)
	}

	getRectangle := func(v1, v2 Loc) rectangle {
		left := min(v1.col, v2.col)
		right := max(v1.col, v2.col)
		top := min(v1.row, v2.row)
		bottom := max(v1.row, v2.row)

		return rectangle{left, right, top, bottom}
	}

	checkEdgeCollisions := func(v1, v2 Loc) bool {
		for left := range n {
			current := getRectangle(v1, v2)
			right := left + 1
			if right >= n {
				right = 0
			}

			cmpV1 := vertices[left]
			cmpV2 := vertices[right]
			if v1 == cmpV1 || v1 == cmpV2 || v2 == cmpV1 || v2 == cmpV2 {
				continue
			}

			cmp := getRectangle(cmpV1, cmpV2)

			if current.left < cmp.right && current.right > cmp.left && current.top < cmp.bottom && current.bottom > cmp.top {
				return false
			}
		}

		return true
	}

	for left := 0; left < n-1; left++ {
		v1 := vertices[left]
		for right := left + 1; right < n; right++ {
			v2 := vertices[right]
			if possible := checkEdgeCollisions(v1, v2); possible {
				area := int((math.Abs(float64(v1.col)-float64(v2.col)) + 1) * (math.Abs(float64(v1.row)-float64(v2.row)) + 1))
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}
