package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Recipe struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title, omitempty"`
}

type Ingredient struct {
	Name string `json:"name, omitempty"`
}

var recipes []Recipe

func GetRecipesEndpoint(w http.ResponseWriter, req *http.Request) {
	for _, item := range recipes {
		json.NewEncoder(w).Encode(item)
		return
	}
	json.NewEncoder(w).Encode(&Recipe{})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	recipes = append(recipes, Recipe{ID: "1", Title: "New recipe"})
	router.HandleFunc("/recipe", GetRecipesEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
