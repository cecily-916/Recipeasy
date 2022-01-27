package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	// router := mux.NewRouter()
	// user_db := goDotEnvVariable("USER_DB")
	// user_pw := goDotEnvVariable("USER_PW")
	// backend_db := goDotEnvVariable("REACT_APP_BACKEND_DATABASE")

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=recipeasy_development sslmode=disable password=postgres") // user_db, user_pw, backend_db)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Recipe{})
	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&Step{})
	db.AutoMigrate(&Quantity{})
	db.AutoMigrate(&UnitMeasurement{})

	for index := range recipes {
		db.Create(&recipes[index])
	}

	for index := range ingredients {
		db.Create(&ingredients[index])
	}

	for index := range steps {
		db.Create(&steps[index])
	}

	for index := range quantities {
		db.Create(&quantities[index])
	}

	for index := range units {
		db.Create(&units[index])
	}

	// router.HandleFunc("/cars", GetCars).Methods("GET")
	// router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	// router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	// router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	// handler := cors.Default().Handler(router)

	// log.Fatal(http.ListenAndServe("localhost:8080", handler))
}

// func GetCars(w http.ResponseWriter, r *http.Request) {
// 	var cars []Car
// 	db.Find(&cars)
// 	json.NewEncoder(w).Encode(&cars)
// }

// func GetCar(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var car Car
// 	db.First(&car, params["id"])
// 	json.NewEncoder(w).Encode(&car)
// }

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
