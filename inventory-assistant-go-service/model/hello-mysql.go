package model

// Post [...]
type Post struct {
	ID        int    `gorm:"primary_key;column:id;type:int;not null"`
	Title     string `gorm:"column:title;type:varchar(191);not null"`
	Content   string `gorm:"column:content;type:varchar(191)"`
	Published int8   `gorm:"column:published;type:tinyint(1);default:0"`
	AuthorID  int    `gorm:"index;column:authorId;type:int"`
}

// User [...]
type User struct {
	ID    int    `gorm:"primary_key;column:id;type:int;not null"`
	Email string `gorm:"unique;column:email;type:varchar(191);not null"`
	Name  string `gorm:"column:name;type:varchar(191)"`
}
