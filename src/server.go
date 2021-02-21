package main

import (
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
	Title string `gorm:"size:10" json:"title"`
	UserID string `gorm:"ForeignKey:id"`
	User User `json:"author"`
}



func main() {
	
	db, _ := database.Connect()
	db.Migrator().DropTable(&User{}, &Post{})
	
	err := db.AutoMigrate(&User{}, &Post{})
	// We need to add foreign keys manually.
	// db.Model(&Post{})
	
	
	
	db.Create(&User{
		Username: "rasel",
		Email:    "rasel@gmail.com",
		Password: "123",
	})
	db.Create(&Post{
		Title:  "dummy post",
		UserID: "1",
	})
	
	if err != nil {
		log.Fatal(err)
	}
	var users User
	// db.Preload("Post").Find(&users)
	//
 db.Raw(`
		SELECT * FROM posts
			join users
		on users.id = posts.user_id
		FOR JSON AUTO
`).Scan(&users)
	

	fmt.Println(users)
	
	// var allPosts Post
	// db.Preload("User").First(&allPosts)
	// st, _:=json.Marshal(allPosts)
	// fmt.Println(string(st))

}
