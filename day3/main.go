package main

import (
	"bufio"
	"fmt"
	"os"
)

func SolveFirstPart(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	next := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}

		return scanner.Text(), true
	}

	var (
		count int
		y     int
	)

	for row, hasNext := next(); hasNext; row, hasNext = next() {
		currPos := (3 * y) % len(row)
		if pos := row[currPos]; string(pos) == "#" {
			count++
		}
		y++
	}

	return count, nil
}

func SolveSecondPart(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var landMap []string

	for scanner.Scan() {
		landMap = append(landMap, scanner.Text())
	}

	var (
		slopes = [][]int{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
	)

	calculateTrees := func(right, down int) int {
		var count int
		var x int
		for y := down; y < len(landMap); y += down {
			x = (x + right) % len(landMap[y])
			if pos := landMap[y][x]; string(pos) == "#" {
				count++
			}
		}

		return count
	}

	var result int
	for _, slope := range slopes {
		trees := calculateTrees(slope[0], slope[1])
		if result == 0 {
			result = trees
			continue
		}

		result = result * trees
	}

	return result, nil
}

func main() {
	fmt.Println(SolveFirstPart("./input.txt"))
	fmt.Println(SolveSecondPart("./input.txt"))
}
