package blue

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Daniel-42-z/lingo-tools/dictutils"
	"github.com/spf13/pflag"
)

func IsSubWord(short, long string, continuous bool) bool {
	if continuous {
		return strings.Contains(long, short)
	}
	longChars := []rune(long)
	shortChars := []rune(short)
	shortLen := len(shortChars)
	if shortLen == 0 {
		return true
	}
	shortIndex := 0
	for _, char := range longChars {
		if char == shortChars[shortIndex] {
			shortIndex++
		}
		if shortIndex == shortLen {
			return true
		}
	}
	return false
}

func MidBlueFindAll(wl dictutils.WordList, q string, continuous bool, action func(string)) {
	for _, word := range wl {
		if IsSubWord(q, word, continuous) {
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

func RunArgs(args []string, dictPath string) error {
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

	MidBlueFindAll(wl, question, continuous, action)
	return nil
}
