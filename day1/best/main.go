package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func solvePartOne(seek int, next func() (int, error)) (int, error) {
	var (
		have   = make(map[int]struct{})
		err    error
		number int
	)
	for ; err == nil; number, err = next() {
		if _, ok := have[seek-number]; ok {
			return number * (seek - number), nil
		}
		have[number] = struct{}{}
	}

	return -1, err
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	fmt.Println(solvePartOne(2020, func() (int, error) {
		if !scanner.Scan() {
			return -1, errors.New("no more rows")
		}

		var n int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &n)

		return n, err
	}))
}
