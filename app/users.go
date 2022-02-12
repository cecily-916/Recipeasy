package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// func getUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	var user User

// 	db.First(&user, params["userid"])

// 	json.NewEncoder(w).Encode(&user)

// 	if err != nil {
// 		return
// 	}
// }

// // WEBHOOKS
// type webhook struct {
// 	Action    string `json:"action"`
// 	Confirmed bool   `json:"isConfirmed"`
// 	Record    struct {
// 		ID       string `json:"userId"`
// 		UserUuid string `json:"userUuid"`
// 		Email    string `json:"email"`
// 		Name     string `json:"name"`
// 	}
// }

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payloadSecret := r.Header.Get("Authorization")
	webhookApiKey := os.Getenv("WEBHOOK_API_KEY")

	if payloadSecret != webhookApiKey {
		fmt.Printf("unauthorized webhook")
		return
	}

	webhookData := make(map[string]interface{})

	fmt.Println(webhookData)
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}

	var user User

	if webhookData["action"] == "create" {
		db.Create(&user)
		db.Save(&user)
	}

	if webhookData["action"] == "update" {
		db.Model(&user).Where("userUuid = ?", webhookData["userUuid"]).First(&user)

		db.Model(&user).UpdateColumn("username & email & name", webhookData["username"], webhookData["email"], webhookData["name"])
	}

	if webhookData["action"] == "delete" {
		db.Model(&user).Where("userUuid = ?", webhookData["userUuid"]).First(&user)
		db.Delete(&user)
	}
}
