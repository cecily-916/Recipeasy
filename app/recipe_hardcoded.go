package main

// var sampleUser = User{
// 	Name:     "Cecily",
// 	Username: "Me!",
// 	Password: "thispassword",
// 	Email:    "email@email.com",
// 	OwnRecipes: []Recipe{
// 		{Title: "Oatmeal",
// 			Description: "This is a delicious breakfast",
// 			Rating:      1,
// 			PrepTime:    "10 mins",
// 			CookTime:    "15 mins",
// 			MainImage:   "https://www.veggieinspired.com/wp-content/uploads/2015/05/healthy-oatmeal-berries-featured.jpg",
// 			Steps: []Step{
// 				{StepOrder: 1,
// 					Details:   "Add oats to the bowl",
// 					Completed: true,
// 					Ingredients: []Ingredient{
// 						{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
// 						{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
// 					},
// 				},
// 				{StepOrder: 2,
// 					Details:   "This is another step",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 1, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
// 					},
// 				},
// 			}},
// 		{Title: "Tea",
// 			Description: "I love tea!",
// 			Rating:      5,
// 			PrepTime:    "2 mins",
// 			CookTime:    "4 mins",
// 			MainImage:   "https://images.pexels.com/photos/1417945/pexels-photo-1417945.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
// 			Steps: []Step{
// 				{
// 					StepOrder: 1,
// 					Details:   "Heat up your water",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 3, Ingredient: "sugar", UnitMeasurement: "cup", IngredientOrder: 1},
// 						{Amount: 10, Ingredient: "salt", UnitMeasurement: "teaspoon", IngredientOrder: 2},
// 					},
// 				},
// 				{
// 					StepOrder: 2,
// 					Details:   "Add tea bag to the water",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
// 					},
// 				},
// 				{
// 					StepOrder: 3,
// 					Details:   "Let steep for 3 mins, then serve!",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 25, Ingredient: "sugar", UnitMeasurement: "pound", IngredientOrder: 1},
// 					},
// 				},
// 			},
// 		},
// 	},
// 	SharedRecipes: []Recipe{
// 		{Title: "Thai Oh My Salmon",
// 			Description: "Salmon over broccoli",
// 			Rating:      5,
// 			PrepTime:    "10 mins",
// 			CookTime:    "20 mins",
// 			Servings:    2,
// 			MainImage:   "https://d2gtpjxvvd720b.cloudfront.net/system/recipe/image/2123/retina_hungry-girl-healthy-thai-oh-my-recipe2-20190528-1730-4972-1885.jpg",
// 			Steps: []Step{
// 				{
// 					StepOrder: 1,
// 					Details:   "Preheat oven to 375 degrees. Lay a large piece of heavy-duty foil on a baking sheet, and spray with nonstick spray.",
// 					Completed: false,
// 				},
// 				{
// 					StepOrder: 2,
// 					Details:   "Add tea bag to the wateIn a small bowl, combine chili sauce, soy sauce, and 2 tsp. water. Mix until uniform.",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 1, Ingredient: "soy sauce", UnitMeasurement: "tbsp", IngredientOrder: 1},
// 						{Amount: 3, Ingredient: "sweet Asian chili sauce", UnitMeasurement: "tbsp", IngredientOrder: 2},
// 						{Amount: 2, Ingredient: "water", UnitMeasurement: "tsp", IngredientOrder: 3},
// 					},
// 				},
// 				{
// 					StepOrder: 3,
// 					Details:   "Distribute broccoli onto the center of the foil. Top with salmon, and sprinkle with 1/4 tsp. garlic powder and 1/8 tsp. paprika. Drizzle with sauce mixture",
// 					Completed: false,
// 					Ingredients: []Ingredient{
// 						{Amount: 2.5, Ingredient: "broccoli florets", UnitMeasurement: "cups", IngredientOrder: 1},
// 						{Amount: 8, Ingredient: "raw skinless salmon fillets (2)", UnitMeasurement: "oz", IngredientOrder: 2},
// 						{Amount: .25, Ingredient: "garlic powder", UnitMeasurement: "tsp", IngredientOrder: 3},
// 						{Amount: .125, Ingredient: "paprika", UnitMeasurement: "tsp", IngredientOrder: 4},
// 					},
// 				},
// 				{
// 					StepOrder: 4,
// 					Details:   "Cover with another large piece of foil. Fold together and seal all four edges of the foil pieces, forming a well-sealed packet.",
// 					Completed: false,
// 				},
// 				{
// 					StepOrder: 5,
// 					Details:   "Bake for 20 minutes, or until salmon is cooked through and veggies are tender.",
// 					Completed: false,
// 				},
// 				{
// 					StepOrder: 6,
// 					Details:   "Cut packet to release hot steam before opening entirely",
// 					Completed: false,
// 				},
// 			},
// 		},
// 	},
// }
