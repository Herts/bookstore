package controllers

import (
	"errors"
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
)

type CartController struct {
	beego.Controller
}

// @Param skuId query string true "skuId"
// @Param amount query int true "amount"
// @router /add [post]
// @Success 200 response response
func (c *CartController) AddToCart(skuId string, amount int) {
	book := models.GetBookBySkuId(skuId)
	if book.ID == 0 {
		c.Data["json"] = response{Message: "SkuId does not exist."}
		c.ServeJSON()
		return
	}
	if amount > book.Amount {
		c.Data["json"] = response{Message: "Your requested amount is above our stock."}
		c.ServeJSON()
		return
	}

	err := models.AddToCart(c.GetSession("uid").(string), skuId, amount)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
	} else {
		c.Data["json"] = response{Message: "Success"}
	}
	c.ServeJSON()
}

// @Param skuId query string true "skuId"
// @Param amount query int true "amount"
// @router /update [put]
func (c *CartController) UpdateCart(skuId string, amount int) {
	book := models.GetBookBySkuId(skuId)
	if book.ID == 0 {
		c.Data["json"] = response{Message: "SkuId does not exist."}
		c.ServeJSON()
		return
	}
	if amount > book.Amount {
		c.Data["json"] = response{Message: "Your requested amount is above our stock."}
		c.ServeJSON()
		return
	}
	var err error
	if amount == 0 {
		err = models.DeleteBookFromCart(c.GetSession("uid").(string), skuId)
	} else if amount > 0 {
		err = models.UpdateBooksInCart(c.GetSession("uid").(string), skuId, amount)
	} else {
		err = errors.New("amount cannot be negative")
	}
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	c.Data["json"] = response{Message: "Success"}
	c.ServeJSON()
}

// @router /placeorder [post]
func (c *CartController) PlaceOrder() {
	c.Data["json"] = response{Message: "Success"}
	err := models.PlaceOrder(c.GetSession("uid").(string))
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
	}
	c.ServeJSON()
}
