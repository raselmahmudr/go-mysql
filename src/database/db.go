package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error){
	f :=  "root:12345@tcp(127.0.0.1:3306)/mydb"
	db, err := gorm.Open(mysql.Open(f), &gorm.Config{} )
	if err != nil {
		fmt.Println(err)
	}
	
	return db, nil
}

