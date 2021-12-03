package main

import (
	"fmt"
	"io"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Run(args []string) int {
	_, err := fmt.Fprint(c.outStream, "Hello, world!\n")

	if err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}
