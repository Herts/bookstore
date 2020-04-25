package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
)

type BookController struct {
	beego.Controller
}

type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// @Param body body models.Book true "Book"
// @Success 200 {response} response
// @router /add [post]
func (c *BookController) CreateOrUpdate() {
	var newBook models.Book
	logs.Debug(newBook)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newBook)
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal" + err.Error()}
		c.ServeJSON()
		return
	}
	logs.Debug(newBook)

	uid := models.GetUserByUuid(c.GetSession("uid").(string)).ID
	b := models.GetOrInitBook(&newBook)
	if b.SkuId == "" {
		b.SkuId = uuid.New().String()
	}
	if b.UserID != uid && b.UserID != 0 {
		c.Data["json"] = response{Message: "This book does belong to you."}
		c.ServeJSON()
		return
	}
	b.UserID = uid
	models.SaveBook(b)
	c.Data["json"] = response{Message: fmt.Sprintf("Book %s saved successfully.", b.Name)}
	c.ServeJSON()
}

// @Success 200 {response} response
// @router /list [get]
func (c *BookController) ListBooks() {
	var newBook models.Book
	logs.Debug(newBook)
}
