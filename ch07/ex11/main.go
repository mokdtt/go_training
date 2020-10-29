package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.deleteItem)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var mu sync.Mutex

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no item given")
		return
	}
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no price given")
		return
	}
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s already exist", item)
		return
	}
	mu.Lock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: created\n", item)
	mu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no item given")
		return
	}
	priceStr := req.URL.Query().Get("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no price given")
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s not exist", item)
		return
	}
	mu.Lock()
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s updated\n", item)
	mu.Unlock()
}

func (db database) deleteItem(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no item given")
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s not exist", item)
		return
	}
	mu.Lock()
	delete(db, item)
	fmt.Fprintf(w, "%s deleted\n", item)
	mu.Unlock()
}
