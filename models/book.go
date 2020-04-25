package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model  `json:"-"`
	UserID      uint
	SkuId       string  `json:"skuId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"decimal"`
	Amount      int     `json:"amount"`
	PicturePath string  `json:"picturePath"`
	IconPath    string  `json:"iconPath"`
}

func GetOrInitBook(b *Book) *Book {
	book := &Book{}
	db.Where(&Book{SkuId: b.SkuId}).FirstOrInit(book, Book{
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
		Amount:      b.Amount,
		PicturePath: b.PicturePath,
		IconPath:    b.IconPath,
	})
	return book
}

func SaveBook(b *Book) {
	db.Save(b)
}

func GetAllBooks() (books []*Book) {
	db.Find(&books)
	return
}

func GetBooksByUser(userId string) (books []*Book) {
	u := User{}
	db.Where(&User{Uuid: userId}).First(&u)
	db.Model(&u).Related(&books, "Book")
	return
}

func GetBookBySkuId(skuId string) *Book {
	b := Book{}
	db.Where(&Book{SkuId: skuId}).First(&b)
	return &b
}
