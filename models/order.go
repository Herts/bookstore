package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderId     string
	OrderDate   time.Time
	OrderItems  []BookWithAmount
	OrderAmount float64 `gorm:"decimal"`
	OrderStatus string
	OrderPaid   bool
	UserID      uint
}

func GetAllOrdersByUser(uuid string) (orders []*Order) {
	u := User{}
	db.Where(&User{Uuid: uuid}).First(&u)
	db.Preload("OrderItems.Book").Model(&u).Related(&orders)
	return
}
