package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserID         uint
	AddressId      string
	RecipientName  string `json:"recipientName"`
	RecipientAddr  string `json:"recipientAddr"`
	RecipientPhone string `json:"recipientPhone"`
}

func GetAllAddressByUser(userId string) (addresses []*Address) {
	u := GetUserByUuid(userId)
	db.Model(&u).Related(&addresses, "Addresses")
	return
}

func SaveAddress(addr *Address) {
	db.Save(&addr)
}

func GetAddressByAddressId(addressId string) *Address {
	addr := Address{}
	db.Where(&Address{AddressId: addressId}).First(&addr)
	return &addr
}

func GetOrInitAddress(addr *Address) *Address {
	address := &Address{}
	db.Where(&Address{AddressId: addr.AddressId}).FirstOrInit(address, Address{
		RecipientName:  addr.RecipientName,
		RecipientAddr:  addr.RecipientAddr,
		RecipientPhone: addr.RecipientPhone,
	})
	return address
}
