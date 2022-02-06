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
	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Deck{})

	// db.Create(&sampleUser)
	// db.Save(&sampleUser)
	// db.Create(&sampleRecipe)
	// db.Save(&sampleRecipe)
	// db.Create(&sampleRecipeTwo)
	// db.Save(&sampleRecipeTwo)
	// db.Create(&s3)
	// db.Save(&s3)

	// router.HandleFunc("/userwebhook", handleUser).Methods("POST")

	router.HandleFunc("/{userid}", getUser).Methods("GET")
	router.HandleFunc("/{userid}/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", handleRecipe).Methods("GET", "DELETE")
	// router.HandleFunc("/user/archive", getArchive).Methods("GET")
	// router.HandleFunc("/user/archive/{id}", handleArchivedRecipe).Methods("PATCH", "DELETE")

	router.HandleFunc("/recipes/{id}/steps", handleStepsByRecipe).Methods("GET", "DELETE")
	router.HandleFunc("/recipes/{id}/steps", createStep).Methods("POST")
	// router.HandleFunc("/recipes/{id}/steps/{stepID}", getStep).Methods("GET")
	// router.HandleFunc("/recipes/{id}/steps/{stepID}", editStep).Methods("PATCH")
	router.HandleFunc("/recipes/{id}/steps/{stepID}", deleteStep).Methods("DELETE")
	router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients", handleIngredientsByStep).Methods("GET", "DELETE")
	router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients", createIngredient).Methods("POST")
	router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients/{ingredientID}", getIngredient).Methods("GET")
	router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients/{ingredientID}", deleteIngredient).Methods("DELETE")
	// router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients/{ingredientID}", editIngredient).Methods("PATCH")

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
