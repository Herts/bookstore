package controllers

import (
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
)

type BuyerController struct {
	beego.Controller
}

func (c *BuyerController) AllOrdersPage() {
	c.Data["Title"] = "All Orders"
	c.Layout = "layout.html"
	c.TplName = "orders.html"
}

func (c *BuyerController) OrderPage() {
	c.Data["Title"] = "Order"
	c.Layout = "layout.html"
	c.TplName = "order.html"
}

func (c *BuyerController) CartPage() {
	userId := c.GetSession("uid")
	if userId != nil {
		c.Data["Title"] = "Cart"
		c.Layout = "layout.html"
		c.TplName = "cart.html"
		cart := models.GetCartByUser(userId.(string))
		c.Data["books"] = cart.Books
	} else {
		c.Redirect("/login", 307)
	}

}

func (c *BuyerController) AllAddressesPage() {
	c.Data["Title"] = "All Address"
	c.Layout = "layout.html"
	c.TplName = "addresses.html"
}

func (c *BuyerController) EditAddressPage() {
	c.Data["Title"] = "Edit Address"
	c.Layout = "layout.html"
	c.TplName = "editaddress.html"
}
