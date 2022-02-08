package main

import (
	"encoding/json"
	"net/http"
)

// handleCollections handles both POST and GET requests
// this request is used to populate the drop down menu of collections and display the associated recipes to the collection

func handleCollections(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var newCollection Collection

		_ = json.NewDecoder(r.Body).Decode(&newCollection)

		db.Create(&newCollection)
		db.Save(&newCollection)
		json.NewEncoder(w).Encode(&newCollection)
	}
}
