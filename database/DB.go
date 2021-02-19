package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hello/models"
)


func DB() *sql.DB  {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		fmt.Println(err)
	}
	return db
}


func DB_INIT()  {
	db := DB()
	statement, _ :=db.Prepare(`CREATE TABLE IF NOT EXISTS users (
	    id integer PRIMARY KEY,
	    name TEXT
	    
	) `)
	
	statement.Exec()
	
	statement, _ = db.Prepare(`INSERT INTO users(name) VALUES (?)`)
	statement.Exec("Alex")
	
	rows, _ :=db.Query(`SELECT * FROM users`)
	
	var users []models.User
	
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Username)
		users = append(users, user)
	}
	
	
}

