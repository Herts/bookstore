package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string    `json:"username" form:"username"`
	Password  string    `json:"password" form:"password"`
	Uuid      string    `json:"userId"`
	Books     []Book    `gorm:"foreignkey:UserID"`
	Orders    []Order   `gorm:"foreignkey:UserID"`
	Addresses []Address `gorm:"foreignkey:UserID"`
	Cart      Cart      `gorm:"foreignkey:UserID"`
}

func GetUserByUserName(username string) *User {
	u := User{}
	db.Where(&User{Username: username}).First(&u)
	return &u
}

func SaveUser(u *User) {
	db.Save(u)
}

func GetUserByUuid(userId string) *User {
	u := User{}
	db.Where(&User{Uuid: userId}).First(&u)
	return &u
}
