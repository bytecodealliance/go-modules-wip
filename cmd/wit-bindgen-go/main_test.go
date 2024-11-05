package main

import (
	"bytes"
	"context"
	"testing"
)

// TestSimpleGenVerbosity ensures that a basic generation case honors the verbose flag
func TestSimpleGenVerbosity(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := Command
	cmd.Writer = &stdout
	cmd.ErrWriter = &stderr

	args := []string{
		"wit-bindgen-go",
		"generate",
		"--world",
		"http-fetch-simple",
		"--dry-run",
		"../../testdata/codegen/simple-http.wit",
	}

	// Run the app without verbose
	cmd.Run(context.Background(), args)
	if stderr.Len() != 0 {
		t.Errorf("output was written to stderr despite lack of --verbose")
	}

	stdout.Reset()
	stderr.Reset()

	cmd.Run(context.Background(), append(args, "-v"))
	if stderr.Len() == 0 {
		t.Errorf("no output was written to stderr when -v was used")
	}

	cmd.Run(context.Background(), append(args, "-vv"))
	if stderr.Len() == 0 {
		t.Errorf("no output was written to stderr when -vv was used")
	}

	cmd.Run(context.Background(), append(args, "--verbose"))
	if stderr.Len() == 0 {
		t.Errorf("no output was written to stderr when --verbose was used")
	}
}