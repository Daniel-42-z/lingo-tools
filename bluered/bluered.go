package bluered

import (
	"errors"
	"fmt"

	"github.com/Daniel-42-z/lingo-tools/dictutils"
	"github.com/spf13/cobra"
	"github.com/Daniel-42-z/lingo-tools/wordutils"
)

type Color int

const (
	Blue Color = iota
	Red
)

func BlueRedFindAll(wl dictutils.WordList, color Color, q string, continuous bool, action func(string)) {
	var cond func(string, string, bool) bool
	if color == Blue {
		cond = func(q, word string, continuous bool) bool { return wordutils.IsSubWord(q, word, continuous) }
	} else {
		cond = func(q, word string, continuous bool) bool { return wordutils.IsSubWord(word, q, continuous) }
	}
	for _, word := range wl {
		if cond(q, word, continuous) {
			action(word)
		}
	}
}

func filterLengthAndPrint(l int) func(string) {
	return func(s string) {
		if len(s) == l {
			fmt.Println(s)
		}
	}
}

func NewBlueRedCmd(color Color, dictPath *string) *cobra.Command {
	name := "blue"
	if color == Red {
		name = "red"
	}

	var (
		filterLength int
		question     string
		continuous   bool
	)

	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Find words for the %s game mode", name),
		RunE: func(cmd *cobra.Command, args []string) error {
			if question == "" {
				return errors.New("Question word not specified. (Specify with -q)")
			}

			var action func(string)
			if filterLength == 0 {
				action = func(s string) { fmt.Println(s) }
			} else {
				action = filterLengthAndPrint(filterLength)
			}

			wl, err := dictutils.MakeWordList(*dictPath)
			if err != nil {
				return fmt.Errorf("failed to load word list: %w", err)
			}

			BlueRedFindAll(wl, color, question, continuous, action)
			return nil
		},
	}

	cmd.Flags().IntVarP(&filterLength, "length", "l", 0, "Only print words of this length (use 0 to not filter)")
	cmd.Flags().StringVarP(&question, "question", "q", "", "Question word")
	cmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Whether the word must be continuous")

	return cmd
}
