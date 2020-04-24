package controllers

import (
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
)

type SellerController struct {
	beego.Controller
}

func (c *SellerController) EditBookPage() {
	c.Data["Title"] = "Edit book"
	c.Layout = "layout.html"
	c.TplName = "editbook.html"
	skuId := c.GetString("skuId")
	if skuId != "" {
		c.Data["edit"] = true
		c.Data["book"] = models.GetBookBySkuId(skuId)
	}

}

func (c *SellerController) AddBookPage() {
	c.Data["Title"] = "Add book"
	c.Layout = "layout.html"
	c.TplName = "editbook.html"
}

func (c *SellerController) AllBooksPage() {
	userId := c.GetSession("uid")
	if userId != nil {
		c.Data["Title"] = "My Books"
		c.Layout = "layout.html"
		c.TplName = "books.html"
		c.Data["manage"] = true
		sellerBooks := models.GetBooksByUser(userId.(string))
		c.Data["books"] = sellerBooks
	} else {
		c.Redirect("/login", 307)
	}

}

func (c *SellerController) ViewBookPage() {
	bookSkuId := c.GetString("skuId")
	if bookSkuId != "" {
		c.Data["Title"] = "Book"
		c.Layout = "layout.html"
		c.TplName = "book.html"
		book := models.GetBookBySkuId(bookSkuId)
		c.Data["book"] = book
		c.Data["manage"] = true
	} else {
		c.Redirect("/", 404)
	}

}
