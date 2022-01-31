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

	sampleRecipe := Recipe{
		Title:       "Oatmeal",
		Description: "This is a delicious breakfast",
		Rating:      1,
		PrepTime:    "10 mins",
		CookTime:    "15 mins",
		MainImage:   "https://www.veggieinspired.com/wp-content/uploads/2015/05/healthy-oatmeal-berries-featured.jpg",
		Steps: []Step{
			{StepOrder: 1,
				Details:   "Add oats to the bowl",
				Completed: true,
				Ingredients: []Ingredient{
					{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
					{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
				},
			},
			{StepOrder: 2,
				Details:   "This is another step",
				Completed: false,
				Ingredients: []Ingredient{
					{Amount: 1, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
				},
			},
		},
	}

	sampleRecipeTwo := Recipe{
		Title:       "Tea",
		Description: "I love tea!",
		Rating:      5,
		PrepTime:    "2 mins",
		CookTime:    "4 mins",
		MainImage:   "https://images.pexels.com/photos/1417945/pexels-photo-1417945.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
		Steps: []Step{
			{
				StepOrder: 1,
				Details:   "Heat up your water",
				Completed: false,
				Ingredients: []Ingredient{
					{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
					{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
				},
			},
			{
				StepOrder: 2,
				Details:   "Add tea bag to the water",
				Completed: false,
				Ingredients: []Ingredient{
					{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
				},
			},
			{
				StepOrder: 3,
				Details:   "Let steep for 3 mins, then serve!",
				Completed: false,
				Ingredients: []Ingredient{
					{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
				},
			},
		},
	}
	db.Create(&sampleRecipe)
	db.Save(&sampleRecipe)
	db.Create(&sampleRecipeTwo)
	db.Save(&sampleRecipeTwo)

	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	// router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", getFullRecipe).Methods("GET")
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
