package main

type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

type Post struct {
	ID      uint   `gorm:"primary_key"`
	Title   string `gorm:"not null"`
	Content string `sql:"type:text;gorm:"not null""`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  uint   `gorm:"not null"`
}
