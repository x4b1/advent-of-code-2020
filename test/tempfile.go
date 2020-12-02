package test

import (
	"io/ioutil"
	"os"
	"testing"
)

func TempFile(t *testing.T) *os.File {
	t.Helper()

	dir := t.TempDir()

	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		t.Fatalf("creating temp file %s", err)
	}
	t.Cleanup(func() {
		err := f.Close()
		if err != nil {
			t.Errorf("closing temp file: %s", err)
		}
	})

	return f
}
