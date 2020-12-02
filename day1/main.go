package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const target = 2020
const filePath = "./input.txt"

func main() {
	numbs, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	n2, _ := multipliedNums(numbs, 0, 2, target)
	fmt.Printf("result for 2 numbers: %d\n", n2)

	n3, _ := multipliedNums(numbs, 0, 3, target)
	fmt.Printf("result for 3 numbers: %d\n", n3)
}

func readFile(path string) (result []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	return result, scanner.Err()
}

func multipliedNums(nums []int, currSum int, nEntries int, target int) (int, bool) {
	if nEntries > 1 {
		for i, n := range nums {
			result, found := multipliedNums(nums[i+1:], currSum+n, nEntries-1, target)
			if found {
				return n * result, true
			}
		}
		return 0, false
	}

	for _, n := range nums {
		if currSum+n == target {
			return n, true
		}
	}

	return 0, false
}
