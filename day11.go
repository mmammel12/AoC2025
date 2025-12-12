package main

import "strings"

type device struct {
	inputs  []*device
	outputs []*device
	name    string
}

func getDeviceMap(lines []string) map[string]*device {
	deviceMap := make(map[string]*device)
	for _, line := range lines {
		name := line[:3]
		d, exists := deviceMap[name]
		if !exists {
			d = &device{[]*device{}, []*device{}, name}
			deviceMap[name] = d
		}

		for output := range strings.SplitSeq(line[5:], " ") {
			o, exists := deviceMap[output]
			if !exists {
				o = &device{[]*device{}, []*device{}, output}
				deviceMap[output] = o
			}
			o.inputs = append(o.inputs, d)
			d.outputs = append(d.outputs, o)
		}
	}

	return deviceMap
}

func solvePaths(start, end string, visited map[string]bool, pathMap map[string]int, deviceMap map[string]*device) int {
	if start == end {
		return 1
	}
	if _, exists := visited[start]; exists || start == "out" {
		return 0
	}
	if paths, exists := pathMap[start]; exists {
		return paths
	}

	visited[start] = true
	current := deviceMap[start]
	paths := 0
	for _, out := range current.outputs {
		paths += solvePaths(out.name, end, visited, pathMap, deviceMap)
	}
	delete(visited, start)
	pathMap[start] = paths

	return paths
}

func day11(lines []string, part int) (int, error) {
	deviceMap := getDeviceMap(lines)
	if part == 1 {
		return solvePaths("you", "out", make(map[string]bool), make(map[string]int), deviceMap), nil
	}

	svrToFft := solvePaths("svr", "fft", make(map[string]bool), make(map[string]int), deviceMap)
	fftToDac := solvePaths("fft", "dac", make(map[string]bool), make(map[string]int), deviceMap)
	dacToOut := solvePaths("dac", "out", make(map[string]bool), make(map[string]int), deviceMap)
	svrToDac := solvePaths("svr", "dac", make(map[string]bool), make(map[string]int), deviceMap)
	dacToFft := solvePaths("dac", "fft", make(map[string]bool), make(map[string]int), deviceMap)
	fftToOut := solvePaths("fft", "out", make(map[string]bool), make(map[string]int), deviceMap)

	ans := (svrToFft * fftToDac * dacToOut) + (svrToDac * dacToFft * fftToOut)
	return ans, nil
}
