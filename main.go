package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello/controllers"
	"hello/database"
	"hello/models"
	"net/http"
)

func main() {
	//
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	var data interface{}
	// 	json.NewDecoder(req.Body).Decode(&data)
	//
	//
	// 	d := data.(map[string]interface{})
	//
	//
	// 	newUser := models.User{
	// 		Username: d["username"].(string),
	// 		Email: d["email"].(string),
	// 	}
	//
	// 	json.NewEncoder(res).Encode(newUser)
	//
	// })
	
	db := database.DB()
	u:= models.User{}
	u.Sync(db)
	
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/users", controllers.GetUsers)
	
	
	fmt.Println("Server is running on port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}






