package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbUri := os.Getenv("DB_URI")
	fmt.Println(dbUri)
	router := mux.NewRouter()

	db, err = gorm.Open("postgres", dbUri)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Recipe{})
	db.AutoMigrate(&Step{})
	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Collection{})

	// db.Create(&sampleUser)
	// db.Save(&sampleUser)

	// db.Create(&sampleUser2)
	// db.Save(&sampleUser2)

	// db.Create(&sampleRecipe)
	// db.Save(&sampleRecipe)
	// db.Create(&sampleRecipeTwo)
	// db.Save(&sampleRecipeTwo)
	// db.Create(&s3)
	// db.Save(&s3)

	router.HandleFunc("/userwebhook", handleWebhook).Methods("POST")

	// GET high level data about user
	router.HandleFunc("/{userfrontid}", getUserData).Methods("GET")
	// GET recipes associated with user
	router.HandleFunc("/{userfrontid}/recipes", getRecipes).Methods("GET")
	// CREATE a new recipe - needs to associate with current user
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	// GET and DELETE endpoints to render individual recipe overview and to delete a recipe
	router.HandleFunc("/recipes/{recipeid}", handleRecipe).Methods("GET", "DELETE")
	// POST to add a new collection
	router.HandleFunc("/collections", createCollection).Methods("POST")
	// GET list of collections associated with a user
	router.HandleFunc("/{userfrontid}/collections", getCollectionsPerUser).Methods("GET")
	// GET and DELETE individual collections
	router.HandleFunc("/collections/{collectionid}", handleCollection).Methods("GET", "DELETE")
	// Assign recipe to a collection
	router.HandleFunc("/recipes/{recipeid}/collections/{collectionid}", handleRecipeCollection).Methods("PATCH", "DELETE")
	router.HandleFunc("/recipes/{recipeid}/collections", getCollectionsPerRecipe).Methods("GET")

	// Create individual step - this is used when editing a recipe and a one-off step needs to be added.
	router.HandleFunc("/recipes/{recipeid}/steps", createStep).Methods("POST")
	// // Edit and delete individual step
	// router.HandleFunc("/recipes/{id}/steps/{stepID}", handleStep).Methods("DELETE", "PUT")
	// // Edit and delete individual ingredient
	// router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients/{id}", handleIngredient).Methods("DELETE", "PUT")
	// Create individual ingredient - this is used when editing a recipe and a one-off step needs to be added.
	router.HandleFunc("/recipes/{id}/steps/{stepID}/ingredients", createIngredient).Methods("POST")

	//::::FOR ARCHIVE FUNCTIONALITY:::::
	router.HandleFunc("/{userfrontid}/archive", getArchive).Methods("GET")
	router.HandleFunc("/archive/{recipeid}", handleArchivedRecipe).Methods("PUT", "DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   ([]string{"https://reci-peasy.herokuapp.com", "http://localhost:3000"}),
		AllowedMethods:   ([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}),
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
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
