package main_test

import (
	"testing"

	day3 "github.com/xabi93/advent-of-code/day3"

	"github.com/xabi93/advent-of-code/test"
)

func TestRun(t *testing.T) {
	f := test.TempFile(t)

	_, err := f.WriteString(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	if err != nil {
		t.Error(err)
	}

	t.Run("first part", func(t *testing.T) {
		result, err := day3.SolveFirstPart(f.Name())
		if err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}

		if result != 7 {
			t.Errorf("expected return 7 but it was %d", result)
		}
	})

	t.Run("second part", func(t *testing.T) {
		result, err := day3.SolveSecondPart(f.Name())
		if err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}

		if result != 336 {
			t.Errorf("expected return 336 but it was %d", result)
		}
	})
}
