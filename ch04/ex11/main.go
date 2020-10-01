package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go_training/ch04/ex11/issueUtil"
)

const usage = `issue command usage
	1. issue create USER/REPO -t TITLE -b BODY
	2. issue read USER/REPO NUMBER
	3. issue edit USER/REPO -n NUMBER -t TITLE -b BODY
	4. issue close USER/REPO -n NUMBER`

const usageCreate = `issue create usage
	issue create USER/REPO TITLE BODY`

const usageRead = `issue read usage
	issue read USER/REPO NUMBER`

const usageEdit = `issue edit usage
	issue edit USER/REPO NUMBER TITLE BODY`

const usageClose = `issue close usage
	issue close USER/REPO NUMBER`

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "create":
		if flag.Arg(2) == "" || flag.Arg(3) == "" {
			fmt.Println(usageCreate)
			os.Exit(1)
		}
		repo := checkRepo()
		//title := flag.Arg(2)
		//body := flag.Arg(3)
		title, err := issueUtil.InputText()
		body, err := issueUtil.InputText()
		issue, err := issueUtil.Create(repo, title, body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "issue: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(issue)
		//issueUtil.ShowIssue(issue, repo, issueNo)
	case "read":
		if flag.Arg(2) == "" {
			fmt.Println(usageRead)
			os.Exit(1)
		}
		repo := checkRepo()
		issueNo := getIssueNo(flag.Arg(2))
		issue, err := issueUtil.Read(repo, issueNo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "issue: %v\n", err)
			os.Exit(1)
		}
		issueUtil.ShowIssue(issue, repo, issueNo)
	case "edit":
		if flag.Arg(2) == "" {
			fmt.Println(usageEdit)
			os.Exit(1)
		}
		if flag.Arg(3) == "" || flag.Arg(4) == "" {
			fmt.Println(usageEdit)
			os.Exit(1)
		}
		repo := checkRepo()
		issueNo := getIssueNo(flag.Arg(2))
		title := flag.Arg(3)
		body := flag.Arg(4)
		issue, err := issueUtil.Edit(repo, issueNo, title, body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "issue: %v\n", err)
			os.Exit(1)
		}
		issueUtil.ShowIssue(issue, repo, issueNo)
	case "close":
		if flag.Arg(2) == "" {
			fmt.Println(usageClose)
			os.Exit(1)
		}
		repo := checkRepo()
		issueNo := getIssueNo(flag.Arg(2))
		issue, err := issueUtil.Close(repo, issueNo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "issue: %v\n", err)
			os.Exit(1)
		}
		issueUtil.ShowIssue(issue, repo, issueNo)
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}

func checkRepo() string {
	repo := flag.Arg(1)
	if len(strings.Split(repo, "/")) != 2 {
		fmt.Fprintf(os.Stderr, "issue: %v\n", "USER/REPO is not specified correctly")
		os.Exit(1)
	}
	return repo
}

func getIssueNo(number string) string {
	if number == "" {
		fmt.Fprintf(os.Stderr, "issue: %v\n", "issue number is not specified")
		os.Exit(1)
	}
	if _, err := strconv.Atoi(number); err != nil {
		fmt.Fprintf(os.Stderr, "issue: %v\n", "not number")
		os.Exit(1)
	}
	return number
}
