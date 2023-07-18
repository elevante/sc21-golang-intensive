package utils

import (
	"day04/ex00/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CheckRequest(w http.ResponseWriter, o *models.Order) {
	if o.CandyCount < 0 {
		response := models.InlineResponse400{Error_: "Negative quantity"}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		switch o.CandyType {
		case "CE":
			CheckMoney(10, w, o)
		case "AA":
			CheckMoney(15, w, o)
		case "NT":
			CheckMoney(17, w, o)
		case "DE":
			CheckMoney(21, w, o)
		case "YR":
			CheckMoney(23, w, o)
		default:
			response := models.InlineResponse400{Error_: "There are no such sweets in the vending machine"}
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func CheckMoney(candyMoney int, w http.ResponseWriter, o *models.Order) {
	if o.Money >= o.CandyCount*candyMoney {
		w.WriteHeader(http.StatusCreated)
		change := o.Money - o.CandyCount*candyMoney
		response := models.InlineResponse201{
			Change: change,
			Thanks: "Thank you!",
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Println(err)
		}
	} else {
		amount := candyMoney*o.CandyCount - o.Money
		errorMessage := fmt.Sprintf("You need %d more money!", amount)

		response := models.InlineResponse400{Error_: errorMessage}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusPaymentRequired)
	}
}
