package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	Books []BookWithAmount
}

type BookWithAmount struct {
	gorm.Model
	Book Book
	Amount int
	CartID uint
}
