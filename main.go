package main

import (
	"fmt"
	"os"

	"github.com/Daniel-42-z/lingo-tools/cipher"
)

func run() error {
	args := os.Args
	if len(args) < 2 {
		return fmt.Errorf("no subcommands specified\nAvailable subcommands: cipher")
	}
	switch args[1] {
	case "cipher":
		return cipher.RunArgs(os.Args[2:])
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
