package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 0-20のComicsの情報を取得します
func getComics() (chan Comic, error) {
	max, err := getComicCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("max", max) //本当の最大
	max = 20                //ここでは20で回す

	nworkers := 5
	comics := make(chan Comic, 5*nworkers)
	comicNums := make(chan int, 1*nworkers)
	done := make(chan int, 0)
	for i := 0; i < nworkers; i++ {
		go fetcher(comicNums, comics, done)
	}

	for i := 1; i <= max; i++ {
		comicNums <- i
	}
	close(comicNums)

	for i := 0; i < nworkers; i++ {
		<-done
	}
	close(done)
	close(comics)
	return comics, nil
}

// comicsの情報をmaxでいくつ取得可能かを返します
func getComicCount() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get main page: %s", resp.Status)
	}
	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, nil
}

func fetcher(comicNums chan int, comics chan Comic, done chan int) {
	for n := range comicNums {
		comic, err := getComic(n)
		if err != nil {
			log.Printf("Can't get comic %d: %s", n, err)
			continue
		}
		comics <- comic
	}
	done <- 1
}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}
