// package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func GetDriver(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var recipe Recipe
// 	var cars []Car

// 	db.First(&driver, params["id"])
// 	db.Model(&driver).Related(&cars)
// 	driver.Cars = cars
// 	json.NewEncoder(w).Encode(&driver)
// }
