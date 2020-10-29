package main

import (
	"fmt"
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

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	fmt.Println("Original:")
	sort.Sort(byLength(tracks))
	printTracks(tracks)

	fmt.Println("\nStableでない Sort (Year -> Title):")
	sort.Sort(byYear(tracks))
	sort.Sort(byTitle(tracks))
	printTracks(tracks)

	fmt.Println("\nStable Sort (Year -> Title):")
	sort.Stable(byYear(tracks))
	sort.Stable(byTitle(tracks))
	printTracks(tracks)

	fmt.Println("\nReset:")
	sort.Sort(byLength(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	prevKey := "Year"  //1つ前のkeyを表す変数があるとする
	sortKey := "Title" //現在ソートしたいkeyを表す変数があるとする
	sort.Sort(customSort{tracks, prevKey, sortKey, func(x, y *Track, prevKey, sortKey string) bool {
		switch sortKey {
		case "Title":
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			if prevKey == "Year" && x.Year != y.Year {
				return x.Year < y.Year
			}
			return false
		case "Year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if prevKey == "Title" && x.Title != y.Title {
				return x.Title < y.Title
			}
			return false
		default:
			panic(prevKey)
		}
	}})
	printTracks(tracks)
}

type customSort struct {
	t                []*Track
	prevKey, sortKey string
	less             func(x, y *Track, prevKey, sortKey string) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.prevKey, x.sortKey) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
