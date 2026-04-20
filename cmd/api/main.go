package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	log.Printf("Starting API server on port 4000")
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/create", createNote)
	mux.HandleFunc("/note", showNote)
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("Hello World!"))
}

func createNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error: 405. Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Here will be note creating"))
}

func showNote(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.NotFound(w,r)
		return
	}

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "Displaying a note with an ID %s", id)
}	