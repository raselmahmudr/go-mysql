package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello/src/database"
	"log"
)

type User struct {
	
	ID uint `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type Post struct {
	ID uint `json:"id"`
	Title string `gorm:"size:60" json:"title"`
	AuthorID string `gorm:"foreignKey=id" json:"author_id"`
	Author User `gorm:"references:id" json:"author"`
}


func main() {
	
	db, _ := database.Connect()
	err := db.Debug().AutoMigrate(&User{}, &Post{})
	if err != nil {
		log.Fatal(err)
	}
	var allPosts Post
	db.Model(&Post{}).Select("*").Joins("left join users on users.id = posts.author_id").Scan(&allPosts)
	st, _:=json.Marshal(allPosts)
	fmt.Println(string(st))
	
	
}
