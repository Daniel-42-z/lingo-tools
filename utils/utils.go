package utils

import (
	"os"
	"strings"
)

type WordMap map[string]struct{}

func MakeWordMap(fileName string) (WordMap, error) {
	wordMap := make(WordMap)
	data, err := os.ReadFile(fileName)
	if err != nil {
		return wordMap, err
	}

	for word := range strings.SplitSeq(string(data), "\n") {
		word = strings.TrimSpace(strings.ToLower(word))
		if word != "" {
			wordMap[word] = struct{}{}
		}
	}
	return wordMap, nil
}

type WordList []string

func MakeWordList(filename string) (WordList, error) {
	wordList := WordList{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return wordList, err
	}
	for word := range strings.SplitSeq(string(data), "\n") {
		word = strings.TrimSpace(strings.ToLower(word))
		if word != "" {
			wordList = append(wordList, word)
		}
	}
	return wordList, nil
}
