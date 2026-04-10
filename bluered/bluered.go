package bluered

import (
	"errors"
	"fmt"

	"github.com/Daniel-42-z/lingo-tools/dictutils"
	"github.com/Daniel-42-z/lingo-tools/wordutils"
	"github.com/spf13/pflag"
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

func RunArgs(args []string, dictPath string, color Color) error {
	fs := pflag.NewFlagSet("blue", pflag.ContinueOnError)
	var (
		filterLength int
		question     string
		continuous   bool
	)
	fs.IntVarP(&filterLength, "length", "l", 0, "Only print words of this length (use 0 to not filter)")
	fs.StringVarP(&question, "question", "q", "", "Question word")
	fs.BoolVarP(&continuous, "continuous", "c", false, "Whether the word must be continuous")
	if len(args) == 0 {
		args = []string{"--help"}
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, pflag.ErrHelp) {
			return nil
		}
		return err
	}

	if question == "" {
		return errors.New("Question word not specified. (Specify with -q)")
	}

	var action func(string)
	if filterLength == 0 {
		action = func(s string) { fmt.Println(s) }
	} else {
		action = filterLengthAndPrint(filterLength)
	}

	wl, err := dictutils.MakeWordList(dictPath)
	if err != nil {
		return fmt.Errorf("Failed to load word list: %s", err)
	}

	BlueRedFindAll(wl, color, question, continuous, action)
	return nil
}
