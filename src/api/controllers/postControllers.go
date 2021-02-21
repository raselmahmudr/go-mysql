package controllers

import (
	"encoding/json"
	"fmt"
	"hello/src/database"
	"hello/src/models"
	"net/http"
)

func GetPost(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	
	fmt.Println("Get all Posts")
	var allPosts []models.CreditCard
	db, _ :=  database.Connect()
	
	db.Table("users").Select("id").Joins("JOIN credit_card on credit_card.user_id === user.id").Find(&allPosts)
	
	json.NewEncoder(res).Encode(allPosts)
}


func CreatePost(res http.ResponseWriter, req *http.Request){
	
	var body map[string]interface{}
	json.NewDecoder(req.Body).Decode(&body)
	
	var newPost models.Post
	newPost.Title = body["title"].(string)
	newPost.Content = body["content"].(string)
	newPost.AuthorID = uint64(body["authorId"].(float64))
	
	db, _ :=  database.Connect()
	db.Create(&newPost)
	
}
