package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"BBB", "Alicia Keys", "As I Am", 2007, length("1m00s")},
	{"NNN", "Martin Solveig", "Smash", 2011, length("2m00s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 2000, length("4m37s")},
	{"Go", "Moby", "Moby", 1992, length("5m37s")},
	{"Go", "Moby", "Moby", 2008, length("6m37s")},
	{"Go", "Moby", "Moby", 1996, length("7m37s")},
	{"AAA", "Martin Solveig", "Smash", 2011, length("8m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

var html = template.Must(template.New("track").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=title">title</a></th>
		<th><a href="?sort=artist">artist</a></th>
		<th><a href="?sort=album">album</a></th>
		<th><a href="?sort=year">year</a></th>
		<th><a href="?sort=length">length</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</tr>
{{end}}
</body>
</html>
`))

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "title":
			sort.Stable(byTitle(tracks))
		case "artist":
			sort.Stable(byArtist(tracks))
		case "album":
			sort.Stable(byAlbum(tracks))
		case "year":
			sort.Stable(byYear(tracks))
		case "length":
			sort.Stable(byLength(tracks))
		}
		err := html.Execute(w, tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
