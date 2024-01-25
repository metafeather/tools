package main

import (
	"os"
	"testing"

	"github.com/metafeather/tools/devenv/internal"
)

func TestStdlibWrite(t *testing.T) {
	stdlib := internal.NewStdlib(stdlib)

	tf, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp dir: %s", err)
	}
	// t.Log("Temp dir:", tf)
	defer os.RemoveAll(tf)

	err = stdlib.Write(tf)
	if err != nil {
		t.Error(err)
	}

	t.Log("Stdlib written")
}
