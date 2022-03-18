
package auto

import (
	"hello/src/database"
	"hello/src/models"
	"log"
)

type Migrator interface {
	AutoMigrator(dst ...interface{})
	
	CreateConstraint(dst interface{})
}



func Load()  {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	
	DB, _ := db.DB()
	defer DB.Close()
	
	// err = db.Debug().AutoMigrate(&models.User{}, &models.Post{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	err = db.Migrator().AutoMigrate(&models.User{}, &models.CreditCard{})
	
	if err != nil {
		log.Fatal(err)
	}
	
	db.Migrator().CreateConstraint(&models.User{}, "CreditCards")
	db.Migrator().CreateConstraint(&models.User{}, "fk_users_credit_cards")
	
	// for post, _ := range posts  {
	// 	db.Model(&models.Post{}).Create(post)
	//
	// }
	// for user, _ := range users  {
	// 	db.Model(&models.User{}).Create(user)
	//
	// }
	
	
}