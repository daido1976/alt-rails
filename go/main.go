package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/todos", handleTodos)
	log.Printf("start server http://localhost%s \n", server.Addr)
	server.ListenAndServe()
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s %s", r.Method, r.URL)
	todo1 := Todo{1, "hoge"}
	todo2 := Todo{2, "fuga"}
	todos := []Todo{todo1, todo2}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(&todos)
	if err != nil {
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
