package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model 
	Number string `json:"number"`
	UserId uint `json:"user_id"`
}