package main

import "net/http"

// // users slice to seed user data.
// var users = []User{
// 	{Name: "Cecily", Username: "cecily916", Password: "abc123"},
// 	{Name: "Selena", Username: "selena08", Password: "efg456"},
// }

// // getUsers responds with the list of all users as JSON
// func getUsers(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, users)
// }

// // postUsers creates a new user from the JSON received in the request body
// func postUsers(c *gin.Context) {
// 	var newUser user

// 	if err := c.BindJSON(&newUser); err != nil {
// 		return
// 	}

// 	users = append(users, newUser)
// 	c.IndentedJSON(http.StatusCreated, newUser)
// }

// WEBHOOKS

func handleUser(w http.ResponseWriter, r *http.Request) {

}
