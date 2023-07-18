package main

import (
	"day04/ex00/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/buy_candy", func(w http.ResponseWriter, r *http.Request) {
		handlers.OrderHandlerRequest(w, r)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
