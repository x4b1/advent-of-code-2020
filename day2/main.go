package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var passCheckerRegx = regexp.MustCompile(`(\d+)-(\d+)\s(\w):\s(.+)`)

type policyApplier = func(firstN int, secondN int, target string, pwd string) bool

const (
	InRangePolicyParam    = "range"
	InPositionPolicyParam = "position"
)

func inRangePolicy(min, max int, target, pwd string) bool {
	c := strings.Count(pwd, target)

	return c >= min && c <= max
}

func inPositionPolicy(first, second int, target, pwd string) bool {
	if len(pwd) < first || len(pwd) < second {
		return false
	}

	return (string(pwd[first-1]) == target) != (string(pwd[second-1]) == target)
}

func policyFromString(s string) (policyApplier, error) {
	switch s {
	case InRangePolicyParam:
		return inRangePolicy, nil
	case InPositionPolicyParam:
		return inPositionPolicy, nil
	}

	return nil, errors.New("unknown policy")
}

func Run(args []string, stdout io.Writer) error {
	if len(args) < 3 {
		return errors.New("missing parameters, usage: <policy> <file-path>")
	}

	policy, err := policyFromString(args[1])
	if err != nil {
		return err
	}

	var valids int
	f, err := os.Open(args[2])
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		matches := passCheckerRegx.FindStringSubmatch(scanner.Text())
		if len(matches) != 5 {
			return fmt.Errorf("invalid data structure: %s", scanner.Text())
		}

		min, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("invalid data structure: %w", err)
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("invalid data structure: %w", err)
		}

		if policy(min, max, matches[3], matches[4]) {
			valids++
		}
	}

	fmt.Fprintf(stdout, "%d", valids)

	return nil
}

func main() {
	if err := Run(os.Args, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
