package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := Run(os.Args, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func Run(args []string, stdout io.Writer) error {
	if len(args) < 4 {
		return errors.New("missing parameters, usage: <entries> <target-number> <file-path>")
	}

	nEntries, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid entries: %w", err)
	}

	target, err := strconv.Atoi(args[2])
	if err != nil {
		return fmt.Errorf("invalid target number: %w", err)
	}

	list, err := numberListFromFile(args[3])
	if err != nil {
		return err
	}

	result, found := findAndMultiplyNums(list, 0, nEntries, target)
	if !found {
		return fmt.Errorf("%d not match for %d number for %s file", target, nEntries, args[3])
	}

	fmt.Fprintf(stdout, "%d", result)

	return nil
}

func numberListFromFile(path string) (result []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	return result, scanner.Err()
}

func findAndMultiplyNums(nums []int, acc int, entries int, target int) (int, bool) {
	if entries > 1 {
		for i, n := range nums {
			result, found := findAndMultiplyNums(nums[i+1:], acc+n, entries-1, target)
			if found {
				return n * result, true
			}
		}
		return 0, false
	}

	for _, n := range nums {
		if acc+n == target {
			return n, true
		}
	}

	return 0, false
}
