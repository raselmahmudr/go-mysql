package models

import "gorm.io/gorm"

// type User struct {
// 	ID uint `gorm:"primary_key; auto_increment" json:"id"`
// 	Username string `gorm:"size:20; not null; " json:"username"`
// 	Email string  `gorm:"size:50; not null; " json:"email"`
// 	Password string  `gorm:"size:60; not null" json:"password"`
// 	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
//
// 	// Post Post `gorm:"foreignKey:ID"`
//
// }
//
// func (u *User)  BeforeSave(db *gorm.DB)  error {
// 	// u.Username = "DSFFFFFFFFFF"
// 	return nil
// }


type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:id" json:"credit_cards"`
}