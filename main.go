package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Greeting string `json:"greeting"`
}

func getGreetingHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Greeting: "Hello World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/message", getGreetingHandler)
	http.ListenAndServe(":8011", nil)

}
