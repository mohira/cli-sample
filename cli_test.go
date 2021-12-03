package main

import (
	"bytes"
	"fmt"
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

func TestCLI_Run_version(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	cli := NewCLI(outStream, errStream)

	command := strings.Split("cli -version", " ")
	code := cli.Run(command)

	if code != ExitCodeOK {
		t.Errorf("got %v want %v", code, ExitCodeOK)
	}

	output := errStream.String()
	want := fmt.Sprintf("cli-sample v%s\n", Version)
	if output != want {
		t.Errorf("got %v want %v", output, want)
	}

}
