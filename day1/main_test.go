package main_test

import (
	"bytes"
	"testing"

	day1 "github.com/xabi93/advent-of-code/day1"

	"github.com/xabi93/advent-of-code/test"
)

func TestRun(t *testing.T) {
	f := test.TempFile(t)

	_, err := f.WriteString("1721\n979\n366\n299\n675\n1456")
	if err != nil {
		t.Error(err)
	}

	t.Run("target 2 numbers", func(t *testing.T) {
		var stdout bytes.Buffer

		if err := day1.Run([]string{"", "2", "2020", f.Name()}, &stdout); err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}
		expected := "514579"
		if out := stdout.String(); out != expected {
			t.Errorf("expected %s, but it was %s", expected, out)
		}
	})

	t.Run("target 3 numbers", func(t *testing.T) {
		var stdout bytes.Buffer

		if err := day1.Run([]string{"", "3", "2020", f.Name()}, &stdout); err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}
		expected := "241861950"
		if out := stdout.String(); out != expected {
			t.Errorf("expected %s, but it was %s", expected, out)
		}
	})
}
