package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	db.First(&user, params["userid"])

	json.NewEncoder(w).Encode(&user)

	if err != nil {
		return
	}
}

// // WEBHOOKS

// func handleUser(w http.ResponseWriter, r *http.Request) {

// }
