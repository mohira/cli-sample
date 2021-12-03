package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK         = 0
	ExitCodeError      = 1
	ExitCodeParseError = 2
)

const (
	Version = "0.1.0"
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet("cli-sample", flag.ContinueOnError)
	flags.BoolVar(&version, "version", false, "output version information and exit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseError
	}

	if version {
		_, err := fmt.Fprintf(c.errStream, "cli-sample v%s\n", Version)
		if err != nil {
			return ExitCodeError
		}

		return ExitCodeOK
	}

	_, err := fmt.Fprint(c.outStream, "Hello, world!\n")

	if err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}
