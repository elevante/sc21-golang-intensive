package main

import (
	"ex00/server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/buy_candy", func(w http.ResponseWriter, r *http.Request) {
		handlers.OrderHandlerRequest(w, r)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
