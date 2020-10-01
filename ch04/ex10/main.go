// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		createdAt := classifyCreatedAt(item.CreatedAt)
		fmt.Printf("#%-5d %7.7s %5s %.55s \n",
			item.Number, item.User.Login, createdAt, item.Title)
	}
}

func classifyCreatedAt(t time.Time) string {
	now := time.Now()
	s := "それ以外"
	if t.After(now.AddDate(0, -1, 0)) {
		s = "1月未満"
	} else if t.After(now.AddDate(-1, 0, 0)) {
		s = "1年未満"
	} else {
		s = "1年以上"
	}
	return s
}
