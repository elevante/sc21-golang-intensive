package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func OrderHandlerRequest(w http.ResponseWriter, r *http.Request) {
	var o Order
	switch r.Method {
	case http.MethodPost:
		PostOrder(w, r, &o)
	default:
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
	}
}

func PostOrder(w http.ResponseWriter, r *http.Request, o *Order) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}
	CheckRequest(w, o)
}
