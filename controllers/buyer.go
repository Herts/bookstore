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
	orders := models.GetAllOrdersByUser(c.GetSession("uid").(string))
	c.Data["orders"] = orders
}

func (c *BuyerController) OrderPage() {
	c.Data["Title"] = "Order"
	c.Layout = "layout.html"
	c.TplName = "order.html"
}

func (c *BuyerController) CartPage() {
	userId := c.GetSession("uid")
	if userId != nil {
		uid := userId.(string)
		c.Data["Title"] = "Cart"
		c.Layout = "layout.html"
		c.TplName = "cart.html"
		cart := models.GetCartByUser(uid)
		c.Data["books"] = cart.Books
		c.Data["addresses"] = models.GetAllAddressByUser(uid)
	} else {
		c.Redirect("/login", 307)
	}

}

func (c *BuyerController) AllAddressesPage() {
	c.Data["Title"] = "All Address"
	c.Layout = "layout.html"
	c.TplName = "addresses.html"
	c.Data["addresses"] = models.GetAllAddressByUser(c.GetSession("uid").(string))
}

func (c *BuyerController) EditAddressPage() {
	c.Data["Title"] = "Edit Address"
	c.Layout = "layout.html"
	c.TplName = "editaddress.html"
	addressId := c.GetString("addressId")
	if addressId != "" {
		c.Data["edit"] = true
		c.Data["address"] = models.GetAddressByAddressId(addressId)
	}
}

func (c *BuyerController) AddAddressPage() {
	c.Data["Title"] = "Add Address"
	c.Layout = "layout.html"
	c.TplName = "editaddress.html"
}
