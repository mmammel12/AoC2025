package main

import (
	"container/heap"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day8(lines []string, part int) (int, error) {
	if part == 1 {
		return day8part1(lines)
	}

	return day8part2(lines)
}

type box struct {
	line    string
	x, y, z int
}

type boxPair struct {
	box1     box
	box2     box
	distance float64
}

type maxHeap []boxPair

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(boxPair))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type minHeap [][]box

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return len(h[i]) < len(h[j])
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.([]box))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func dist(a, b box) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2.0) + math.Pow(float64(a.y-b.y), 2.0) + math.Pow(float64(a.z-b.z), 2.0))
}

func getBox(line string) box {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return box{line, x, y, z}
}

func day8part1(lines []string) (int, error) {
	var maxSize int
	if len(lines) == 20 {
		maxSize = 10
	} else {
		maxSize = 1000
	}

	h := maxHeap{}
	heap.Init(&h)

	left := 0
	for left < len(lines)-1 {
		right := left + 1
		leftBox := getBox(lines[left])
		for right < len(lines) {
			rightBox := getBox(lines[right])
			pair := boxPair{leftBox, rightBox, dist(leftBox, rightBox)}
			heap.Push(&h, pair)
			if h.Len() > maxSize {
				heap.Pop(&h)
			}
			right++
		}
		left++
	}

	appendShared := func(m map[string]*[]box, key string, newBoxes []box) {
		s := m[key]
		*s = append(*s, newBoxes...)
	}

	junctions := make(map[string]*[]box)
	for _, pair := range h {
		s1, exists1 := junctions[pair.box1.line]
		s2, exists2 := junctions[pair.box2.line]
		if !exists1 && !exists2 {
			s := &[]box{pair.box1, pair.box2}
			junctions[pair.box1.line] = s
			junctions[pair.box2.line] = s
		} else if !exists1 {
			appendShared(junctions, pair.box2.line, []box{pair.box1})
			s := junctions[pair.box2.line]
			junctions[pair.box1.line] = s
		} else if !exists2 {
			appendShared(junctions, pair.box1.line, []box{pair.box2})
			s := junctions[pair.box1.line]
			junctions[pair.box2.line] = s
		} else if s1 != s2 {
			appendShared(junctions, pair.box1.line, *s2)
			for _, b := range *s2 {
				junctions[b.line] = junctions[pair.box1.line]
			}
		}
	}

	jSlice := make([]*[]box, 0)
	for _, ptr := range junctions {
		if !slices.Contains(jSlice, ptr) {
			jSlice = append(jSlice, ptr)
		}
	}

	jHeap := minHeap{}
	heap.Init(&jHeap)
	for _, ptr := range jSlice {
		heap.Push(&jHeap, *ptr)
		if jHeap.Len() > 3 {
			heap.Pop(&jHeap)
		}
	}

	pass := 1
	for _, junct := range jHeap {
		pass *= len(junct)
	}

	return pass, nil
}

type minHeap2 []boxPair

func (h minHeap2) Len() int { return len(h) }
func (h minHeap2) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}
func (h minHeap2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap2) Push(x any) {
	*h = append(*h, x.(boxPair))
}

func (h *minHeap2) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func day8part2(lines []string) (int, error) {
	h := maxHeap{}
	heap.Init(&h)

	left := 0
	for left < len(lines)-1 {
		right := left + 1
		leftBox := getBox(lines[left])
		for right < len(lines) {
			rightBox := getBox(lines[right])
			pair := boxPair{leftBox, rightBox, dist(leftBox, rightBox)}
			heap.Push(&h, pair)
			right++
		}
		left++
	}

	appendShared := func(m map[string]*[]box, key string, newBoxes []box) {
		s := m[key]
		*s = append(*s, newBoxes...)
	}

	data := make([]boxPair, len(h))
	copy(data, h)
	mh := minHeap2(data)
	heap.Init(&mh)

	pass := 1
	junctions := make(map[string]*[]box)
	for mh.Len() != 0 {
		pair := heap.Pop(&mh).(boxPair)
		s1, exists1 := junctions[pair.box1.line]
		s2, exists2 := junctions[pair.box2.line]
		if !exists1 && !exists2 {
			s := &[]box{pair.box1, pair.box2}
			junctions[pair.box1.line] = s
			junctions[pair.box2.line] = s
		} else if !exists1 {
			appendShared(junctions, pair.box2.line, []box{pair.box1})
			s := junctions[pair.box2.line]
			junctions[pair.box1.line] = s
		} else if !exists2 {
			appendShared(junctions, pair.box1.line, []box{pair.box2})
			s := junctions[pair.box1.line]
			junctions[pair.box2.line] = s
		} else if s1 != s2 {
			appendShared(junctions, pair.box1.line, *s2)
			for _, b := range *s2 {
				junctions[b.line] = junctions[pair.box1.line]
			}
		}

		if s1 != nil && len(*s1) == len(lines) {
			pass *= pair.box1.x * pair.box2.x
			break
		}
		if s2 != nil && len(*s2) == len(lines) {
			pass *= pair.box1.x * pair.box2.x
			break
		}
	}

	return pass, nil
}
