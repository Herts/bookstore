package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
)

type AddressController struct {
	beego.Controller
}

// @Param body body models.Address true "Address"
// @Success 200 {response} response
// @router /add [post]
func (c *AddressController) CreateOrUpdate() {
	var newAddr models.Address
	logs.Debug(newAddr)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newAddr)
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal" + err.Error()}
		c.ServeJSON()
		return
	}
	logs.Debug(newAddr)
	address := models.GetOrInitAddress(&newAddr)
	if address.AddressId == "" {
		address.AddressId = uuid.New().String()
	}

	uid := models.GetUserByUuid(c.GetSession("uid").(string)).ID
	if address.UserID != uid && address.UserID != 0 {
		c.Data["json"] = response{Message: "This address does belong to you."}
		c.ServeJSON()
		return
	}

	address.UserID = uid
	models.SaveAddress(address)
	c.Data["json"] = response{Message: fmt.Sprintf("Address %s saved successfully.", address.RecipientAddr)}
	c.ServeJSON()
}
