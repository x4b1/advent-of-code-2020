package main_test

import (
	"bytes"
	"testing"

	day2 "github.com/xabi93/advent-of-code/day2"

	"github.com/xabi93/advent-of-code/test"
)

func TestRun(t *testing.T) {
	f := test.TempFile(t)

	_, err := f.WriteString(`1-3 a: abcde
	1-3 b: cdefg
	2-9 c: ccccccccc`)
	if err != nil {
		t.Error(err)
	}

	t.Run("in range policy", func(t *testing.T) {
		var stdout bytes.Buffer

		if err := day2.Run([]string{"", day2.InRangePolicyParam, f.Name()}, &stdout); err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}
		expected := "2"
		if out := stdout.String(); out != expected {
			t.Errorf("expected %s, but it was %s", expected, out)
		}
	})

	t.Run("in position policy", func(t *testing.T) {
		var stdout bytes.Buffer

		if err := day2.Run([]string{"", day2.InPositionPolicyParam, f.Name()}, &stdout); err != nil {
			t.Errorf("expected to not return error but it returns %v", err)
			t.FailNow()
		}
		expected := "1"
		if out := stdout.String(); out != expected {
			t.Errorf("expected %s, but it was %s", expected, out)
		}
	})
}
