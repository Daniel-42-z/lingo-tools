package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Daniel-42-z/lingo-tools/blue"
	"github.com/Daniel-42-z/lingo-tools/cipher"
	"github.com/spf13/pflag"
)

func run() error {
	fs := pflag.NewFlagSet("lingo-tools", pflag.ContinueOnError)
	fs.SetInterspersed(false)
	dictPath := fs.StringP("dict", "d", "words.txt", "Path to word list used")

	if err := fs.Parse(os.Args[1:]); err != nil {
		if errors.Is(err, pflag.ErrHelp) {
			return nil
		}
		return err
	}

	args := fs.Args()
	if len(args) < 1 {
		return fmt.Errorf("no subcommands specified\nAvailable subcommands: cipher, blue")
	}

	switch args[0] {
	case "cipher":
		return cipher.RunArgs(args[1:], *dictPath)
	case "blue":
		return blue.RunArgs(args[1:], *dictPath)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
