package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	Books  []BookWithAmount
	UserID uint
}

type BookWithAmount struct {
	gorm.Model
	Book   Book
	SkuId  string
	Amount int
	CartID uint
}

func GetCartByUser(userId string) *Cart {
	c := Cart{}
	u := User{}
	db.Where(&User{Uuid: userId}).Find(&u)
	db.Model(&u).Related(&c)
	db.Model(&c).Related(&c.Books)
	for i := range c.Books {
		db.Where(&Book{SkuId: c.Books[i].SkuId}).First(&c.Books[i].Book)
	}
	return &c
}
