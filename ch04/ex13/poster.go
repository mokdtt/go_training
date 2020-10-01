package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Movie struct {
	Title  string
	Year   string
	Poster string
}

const omdbURL = "https://www.omdbapi.com/?t=%s&apikey=%s"

func getMovie(title string) (*Movie, error) {
	title = url.QueryEscape(title)
	apikey := url.QueryEscape(os.Getenv("OMDB_KEY"))
	queryURL := fmt.Sprintf(omdbURL, title, apikey)
	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var movie Movie
	if err = json.Unmarshal(body, &movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

func getPoster(posterURL string, filename string) error {
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("search query failed: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename += ".jpg"
	err = ioutil.WriteFile(filename, body, 0666)
	if err != nil {
		return err
	}
	return nil
}
