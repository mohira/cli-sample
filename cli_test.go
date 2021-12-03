package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	cli := NewCLI(outStream, errStream)

	command := strings.Split("cli", " ")
	code := cli.Run(command)
	if code != ExitCodeOK {
		t.Errorf("got %v want %v", code, ExitCodeOK)
	}

	output := outStream.String()
	want := "Hello, world!\n"
	if output != want {
		t.Errorf("got %v want %v", output, want)
	}

}
