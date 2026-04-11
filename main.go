package main

import (
	"fmt"
	"os"

	"github.com/Daniel-42-z/lingo-tools/bluered"
	"github.com/Daniel-42-z/lingo-tools/cipher"
	"github.com/spf13/cobra"
)

func main() {
	var dictPath string

	rootCmd := &cobra.Command{
		Use:   "lingo-tools",
		Short: "A collection of tools for word games",
	}

	rootCmd.PersistentFlags().StringVarP(&dictPath, "dict", "d", "words.txt", "Path to word list used")

	rootCmd.AddCommand(cipher.NewCipherCmd(&dictPath))
	rootCmd.AddCommand(bluered.NewBlueRedCmd(bluered.Blue, &dictPath))
	rootCmd.AddCommand(bluered.NewBlueRedCmd(bluered.Red, &dictPath))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
