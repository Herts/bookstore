package controllers

import "github.com/astaxie/beego"

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
	c.Data["Title"] = "Cart"
	c.Layout = "layout.html"
	c.TplName = "cart.html"
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
