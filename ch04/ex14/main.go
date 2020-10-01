package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"go_training/ch04/ex14/githubhtml"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var milestoneList = template.Must(template.New("milestonelist").Parse(`
<h1>milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Creator</th>
  <th>Title</th>
  <th>Description</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
  <td>{{.Title}}</td>
  <td>{{.Description}}</td>
</tr>
{{end}}
</table>
`))

var contributorList = template.Must(template.New("contribuotrlist").Parse(`
<h1>contributors</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>User</th>
  <th>Contributions</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.ID}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
  <td>{{.Contributions}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handlerIssue)
	http.HandleFunc("/milestone", handlerMilestone)
	http.HandleFunc("/contributor", handlerContributor)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handlerIssue(w http.ResponseWriter, r *http.Request) {
	repo := os.Args[1]
	// milestone実行
	result, err := githubhtml.SearchIssues(repo)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

func handlerMilestone(w http.ResponseWriter, r *http.Request) {
	repo := os.Args[1]
	// milestone実行
	result, err := githubhtml.GetMilestone(repo)
	if err != nil {
		log.Fatal(err)
	}
	if err := milestoneList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

func handlerContributor(w http.ResponseWriter, r *http.Request) {
	repo := os.Args[1]
	// contributor実行
	result, err := githubhtml.GetContributor(repo)
	if err != nil {
		log.Fatal(err)
	}
	if err := contributorList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
