package main

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	t.Parallel()

	cmd := newVersionCmd()
	output := &bytes.Buffer{}

	cmd.SetOut(output)
	cmd.SetArgs(nil) // No arguments for the version command

	if err := cmd.Execute(); err != nil {
		t.Errorf("Error executing version command: %v", err)
	}

	wantedOutput := "vigilo version dev\ncommit: none\ndate: unknown\n"
	if output.String() != wantedOutput {
		t.Errorf("Unexpected output from version command. Got: %q, Want: %q", output.String(), wantedOutput)
	}
}
