package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var OOSError = errors.New("requested amount is beyond stock")

type Cart struct {
	gorm.Model
	Books  []BookWithAmount
	UserID uint
}

type BookWithAmount struct {
	gorm.Model
	Book   Book `gorm:"foreignkey:SkuId;association_foreignkey:SkuId;association_autoupdate:false"`
	SkuId  string
	Amount int
	CartID uint
}

func GetCartByUser(userId string) *Cart {
	c := Cart{}
	u := GetUserByUuid(userId)
	db.Preload("Books.Book").Model(&u).Related(&c)
	return &c
}

func AddToCart(userId string, skuId string, amount int) error {
	// Check if book is OOS
	book := GetBookBySkuId(skuId)
	if book.Amount < amount {
		return OOSError
	}

	// Check if cart exists
	c := GetCartByUser(userId)
	if c.ID == 0 {
		c.Books = append(c.Books, BookWithAmount{SkuId: skuId, Amount: amount})
		result := db.Create(&c)
		return result.Error
	}

	// Check if cart contains the same book
	existed := false
	for i, b := range c.Books {
		if b.SkuId == skuId {
			existed = true
			if amount+b.Amount > b.Book.Amount {
				return OOSError
			} else {
				c.Books[i].Amount += amount
			}
		}
	}
	if !existed {
		c.Books = append(c.Books, BookWithAmount{SkuId: skuId, Amount: amount})
	}
	result := db.Save(&c)

	return result.Error
}

func DeleteBookFromCart(userId string, skuId string) error {
	// Check if cart exists
	c := GetCartByUser(userId)
	if c.ID == 0 {
		result := db.Create(&c)
		return result.Error
	}
	for _, b := range c.Books {
		if b.SkuId == skuId {
			result := db.Model(&c).Association("Books").Delete(b)
			return result.Error
		}
	}
	return nil

}

func UpdateBooksInCart(userId string, skuId string, amount int) error {
	if amount == 0 {
		return DeleteBookFromCart(userId, skuId)
	}

	// Check if cart exists
	c := GetCartByUser(userId)
	if c.ID == 0 {
		result := db.Create(&c)
		return result.Error
	}

	for i, b := range c.Books {
		if b.SkuId == skuId {
			c.Books[i].Amount = amount
			result := db.Save(&c)
			return result.Error
		}
	}
	return nil
}
