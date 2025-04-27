package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	var mu sync.Mutex
	store := NewStore()
	router := NewRouter(store)

	http.Handle("/", router)

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		store = NewStore()
		router = NewRouter(store)
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
