package main

import (
		"log"
		"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Printf("Starting API server on port 4000")
	
	mux.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":4000", mux))
}


func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддеживается", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("Привет Мир!"))
}