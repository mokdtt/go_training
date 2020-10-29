package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var dbList = template.Must(template.New("dblist").Parse(`
<h1>Items</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $item, $price := .}}
<tr>
  <td>{{$item}}</td>
  <td>{{$price}}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var tmp bytes.Buffer
	if err := dbList.Execute(&tmp, db); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, tmp.String())
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
