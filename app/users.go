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
type Webhook struct {
	Action    string `json:"action"`
	Confirmed bool   `json:"isConfirmed"`
	Record    struct {
		UserId int `json:"userId"`

		UserUuid string `json:"userUuid"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payloadSecret := r.Header.Get("Authorization")
	webhookApiKey := os.Getenv("WEBHOOK_API_KEY")
	fmt.Println(payloadSecret)
	fmt.Println(webhookApiKey)

	fmt.Println("accessed")
	if payloadSecret != webhookApiKey {
		fmt.Printf("unauthorized webhook")
		return
	}
	var webhook Webhook

	fmt.Println(webhook)
	err := json.NewDecoder(r.Body).Decode(&webhook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload: ", webhook)

	if webhook.Action == "create" {
		user := User{
			UserfrontId: webhook.Record.UserId,
			UserUuid:    webhook.Record.UserUuid,
			Email:       webhook.Record.Email,
			Name:        webhook.Record.Name,
			Username:    webhook.Record.Username,
		}

		db.Create(&user)
		db.Save(&user)
	}

	// if webhookData["action"] == "update" {
	// 	var user User

	// 	db.Model(&user).Where("userUuid = ?", webhookData["userUuid"]).First(&user)

	// 	db.Model(&user).UpdateColumn("username & email & name", webhookData["username"], webhookData["email"], webhookData["name"])
	// }

	// if webhookData["action"] == "delete" {
	// 	var user User

	// 	db.Model(&user).Where("userUuid = ?", webhookData["userUuid"]).First(&user)
	// 	db.Delete(&user)
	// }
}
