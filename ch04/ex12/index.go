package main

import (
	"bufio"
	"encoding/gob"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func indexing(filename string, comicChan chan Comic) error {
	wordIndex, numIndex := indexComics(comicChan)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	//fmt.Println(wordIndex)
	err = enc.Encode(wordIndex)
	if err != nil {
		return err
	}
	err = enc.Encode(numIndex)
	if err != nil {
		return err
	}
	return nil
}

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)
	for comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		scanner.Split(ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
	}
	return wordIndex, numIndex
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	start := 0
	stop := 0
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			start = i - size
			break
		}
	}
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if !unicode.IsLetter(r) {
			stop = i - size
			break
		}
	}
	if stop > start {
		token = data[start:stop]
	}
	return i, token, nil
}
