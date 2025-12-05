package main

import (
	"slices"
	"strconv"
	"strings"
)

func day5(lines []string, part int) (int, error) {
	if part == 1 {
		return day5part1(lines)
	}

	return day5part2(lines)
}

type Range struct {
	start int
	end   int
}

func day5part1(lines []string) (int, error) {
	pass := 0

	isFreshList := true
	freshRanges := []Range{}
	for _, line := range lines {
		if line == "" {
			isFreshList = false
			continue
		}

		if isFreshList {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			freshRanges = append(freshRanges, Range{start: start, end: end})
		} else {
			id, _ := strconv.Atoi(line)
			for _, freshRange := range freshRanges {
				if id >= freshRange.start && id <= freshRange.end {
					pass++
					break
				}
			}
		}
	}

	return pass, nil
}

func day5part2(lines []string) (int, error) {
	pass := 0

	isFreshList := true
	freshRanges := []Range{}
	for _, line := range lines {
		if line == "" {
			break
		}

		if isFreshList {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			freshRanges = append(freshRanges, Range{start: start, end: end})
		}
	}

	slices.SortFunc(freshRanges, func(a, b Range) int {
		return a.start - b.start
	})

	collapsedRanges := []Range{}
	left := 0
	for left < len(freshRanges) {
		right := left + 1
		currentRange := freshRanges[left]
		for right < len(freshRanges) {
			compareRange := freshRanges[right]
			if compareRange.start >= currentRange.start && compareRange.end <= currentRange.end {
				right++
				continue
			}
			if compareRange.start <= currentRange.end+1 && compareRange.end > currentRange.end {
				currentRange.end = compareRange.end
			} else {
				break
			}
			right++
		}
		collapsedRanges = append(collapsedRanges, currentRange)
		left = right
	}

	for _, collapsedRange := range collapsedRanges {
		pass += (collapsedRange.end - collapsedRange.start) + 1
	}

	return pass, nil
}
