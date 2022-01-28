package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

// func goDotEnvVariable(key string) string {
// 	err := godotenv.Load(".env")

// if err != nil {
// 	log.Fatal("Error loading .env file")
// }
// return os.Getenv(key)
// }

func main() {
	router := mux.NewRouter()
	// user_db := goDotEnvVariable("USER_DB")
	// user_pw := goDotEnvVariable("USER_PW")
	// backend_db := goDotEnvVariable("REACT_APP_BACKEND_DATABASE")

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=recipeasy_test sslmode=disable password=postgres") // user_db, user_pw, backend_db)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Recipe{})
	db.AutoMigrate(&Step{})
	db.AutoMigrate(&Quantity{})

	sampleRecipe := Recipe{
		Title: "Oatmeal",
		Steps: []Step{
			{
				Details:   "This is a step",
				Completed: true,
				Quantities: []Quantity{
					{Amount: 3, Ingredient: "sugar", UnitMeasurement: ""},
					{Amount: 10, Ingredient: "salt", UnitMeasurement: ""},
				},
			},
			{
				Details:   "This is another step",
				Completed: false,
				Quantities: []Quantity{
					{Amount: 25, Ingredient: "sugar", UnitMeasurement: ""},
				},
			},
		},
	}

	db.Create(&sampleRecipe)
	db.Save(&sampleRecipe)

	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", getRecipeByID).Methods("GET")
	// router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	// router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	// router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}

// func GetDriver(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var driver Driver
// 	var cars []Car
// 	db.First(&driver, params["id"])
// 	db.Model(&driver).Related(&cars)
// 	driver.Cars = cars
// 	json.NewEncoder(w).Encode(&driver)
// }

// func DeleteCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var car Car
// 	db.First(&car, params["id"])
// 	db.Delete(&car)

// 	var cars []Car
// 	db.Find(&cars)
// 	json.NewEncoder(w).Encode(&cars)
// }
