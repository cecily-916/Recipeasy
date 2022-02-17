# Recipeasy | Recipes Made Easy Backend

API for Recipeasy, a React app created as a student capstone project for the Ada Developer Academy. The website makes working through a recipe easy-peasy by breaking instructions down into clear steps and adding visual aids to the ingredients to help users avoid the dreaded tablespoon versus teaspoon confusion. Recipeasy allows users to input their own recipes, organize them into collections, and add them to their Google calendar to assist in meal planning.

# Dependencies
This API requires the following dependencies:<br>
<br>
Go version: `go1.17.6 darwin/amd64`<br>
Userfront webhook

  `github.com/gorilla/mux v1.8.0`<br>
  `github.com/jinzhu/gorm v1.9.16`<br>
  `github.com/joho/godotenv v1.4.0`<br>
  `github.com/rs/cors v1.8.2`
  
# Environment Variables
- DB_URI="host=%s, port=%s user=%s dbname=%s sslmode=disable password=%s"
- WEBHOOK_API_KEY 

# To-Dos
- Clean up unused endpoints
- Add tests and error handling
- Clean up data models



