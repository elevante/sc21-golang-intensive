package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/buy_candy", func(w http.ResponseWriter, r *http.Request) {
		OrderHandlerRequest(w, r)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
