package controllers

import (
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "Bookstore"
	c.Layout = "layout.html"
	c.TplName = "books.html"
	allBooks := models.GetAllBooks()
	c.Data["books"] = allBooks
}

func (c *MainController) ViewBookPage() {
	bookSkuId := c.GetString("skuId")
	if bookSkuId != "" {
		c.Data["Title"] = "Book"
		c.Layout = "layout.html"
		c.TplName = "book.html"
		book := models.GetBookBySkuId(bookSkuId)
		c.Data["book"] = book
	} else {
		c.Redirect("/", 404)
	}
}

func (c *MainController) ContactPage() {
	c.Layout = "layout.html"
	c.TplName = "contact.html"
	c.Data["Title"] = "Contact"
}
