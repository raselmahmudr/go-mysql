package controllers

import (
	"encoding/json"
	"hello/src/database"
	"hello/src/models"
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("Get all Users"))
	db, _ :=database.Connect()
	var users []models.User
	db.Find(&users)
	
	json.NewEncoder(res).Encode(users)
	
}

func CreateUser(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	
	// db, _ :=database.Connect()
	//
	//  db.Create(&models.User{
	// 	Username:  "ewrwe",
	// 	Email:     "wer",
	// 	Password:  "wer",
	// })
	//

}

