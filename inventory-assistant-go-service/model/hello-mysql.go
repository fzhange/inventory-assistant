package model

import "gorm.io/gorm"

// Post [...]
type Post struct {
	gorm.Model
	Title  string `gorm:"column:title;type:varchar(191);not null"`
	UserID uint
}

// User [...]
type User struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(191)"`
	Post []Post
}
