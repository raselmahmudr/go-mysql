package main

import (
	"database/sql"
	_ "database/sql"
	"log"
)




func dbConnection () *sql.DB {
	dbDriver:= "mysql"
	dbUser:= "root"
	dbPass:= "123"
	dbName:= "mydb"
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName )
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	
	
	// db := database.DB_INIT()
	// db.AutoMigrate(&Product{}, &User{})
	
	// db.Create(&Product{Title: "Jeans Pant",		Price: 23.34})
	// db.Create(&User{Username: "rasel", Email:"rasel@gmail.com",Password: "123"	})
	
	// var user User
	// db.Where("username = ? ", "rasel").First(&user)
	// fmt.Println(user)
	//
	//
	// http.HandleFunc("/register", controllers.Register)
	// http.HandleFunc("/login", controllers.Login)
	// http.HandleFunc("/users", controllers.GetUsers)
	
	
	// fmt.Println("Server is running on port 4000")
	// if err := http.ListenAndServe(":4000", nil); err != nil {
	// 	fmt.Println(err)
	// }
}









