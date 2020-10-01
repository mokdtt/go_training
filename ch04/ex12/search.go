package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

func readIndex(filename string) (WordIndex, NumIndex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	dec := gob.NewDecoder(file)
	var wordIndex WordIndex
	var numIndex NumIndex
	err = dec.Decode(&wordIndex)
	if err != nil {
		return nil, nil, err
	}
	dec.Decode(&numIndex)
	if err != nil {
		return nil, nil, err
	}
	return wordIndex, numIndex, nil
}

func comicsContainingWords(words []string, wordIndex WordIndex, numIndex NumIndex) []Comic {
	found := make(map[int]int) // comic Num -> count words found
	comics := make([]Comic, 0)
	for _, word := range words {
		for num := range wordIndex[word] {
			found[num]++
		}
	}
	for num, nfound := range found {
		if nfound == len(words) {
			comics = append(comics, numIndex[num])
		}
	}
	return comics
}

func searchWord(query string, filename string) error {
	wordIndex, numIndex, err := readIndex(filename)
	if err != nil {
		return err
	}
	comics := comicsContainingWords(strings.Fields(query), wordIndex, numIndex)
	for _, comic := range comics {
		fmt.Printf("%+v\n\n", comic)
	}
	return nil
}
