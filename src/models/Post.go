package models

import "time"

type Post struct {
	ID uint64 `gorm:"primary_key; auto_increment" json:"id"`
	Title string `gorm:"size:30; not null; unique" json:"title"`
	Content string `gorm:"type:text;" json:"content"`
	Author User `gorm:"auto_increment"`
	AuthorID uint64 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}