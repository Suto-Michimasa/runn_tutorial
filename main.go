package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	store *Store
	mu    sync.Mutex
)

func main() {
	store = NewStore()
	router := NewRouter(store)

	http.Handle("/", router)

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		store = NewStore()
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
