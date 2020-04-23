package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	AddressId      string
	RecipientAddr  string
	RecipientPhone string
	UserID         uint
}
