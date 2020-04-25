package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderDate   time.Time
	OrderItems  []BookWithAmount
	OrderAmount float64 `gorm:"decimal"`
	OrderStatus string
	OrderPaid   bool
	UserID      uint
	Address     Address `gorm:"foreignkey:AddressId;association_foreignkey:AddressId;association_autoupdate:false"`
	AddressId   string
}

func GetAllOrdersByUser(uuid string) (orders []*Order) {
	u := User{}
	db.Where(&User{Uuid: uuid}).First(&u)
	db.Preload("OrderItems.Book").Preload("Address").Model(&u).Related(&orders)
	return
}

func PlaceOrder(userId string, addressId string) error {
	c := GetCartByUser(userId)
	if len(c.Books) == 0 {
		return errors.New("no books in the cart")
	}
	var total float64
	for i, b := range c.Books {
		if b.Amount > b.Book.Amount {
			return errors.New(fmt.Sprintf("There is only %d Book %s left.", b.Book.Amount, b.Book.Name))
		}
		total += float64(b.Amount) * b.Book.Price
		c.Books[i].CartID = 0
	}

	order := Order{
		OrderDate:   time.Now(),
		OrderItems:  c.Books,
		OrderAmount: total,
		OrderStatus: "",
		OrderPaid:   false,
		UserID:      c.UserID,
		AddressId:   addressId,
	}
	result := db.Save(&order)
	for _, b := range c.Books {
		b.Book.Amount -= b.Amount
		db.Save(&b.Book)
	}
	return result.Error
}
