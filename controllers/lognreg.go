package controllers

import (
	"github.com/Herts/bookstore/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/google/uuid"
)

type LoginController struct {
	beego.Controller
}

type RegController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
		logs.Error(err)
		c.Redirect("/login", 302)
		c.Data["error"] = err.Error()
		return
	}

	if u.Username == "" || u.Password == "" {
		c.Get()
		c.Data["error"] = "Username / password should not be empty"
		return
	}

	u2 := models.GetUserByUserName(u.Username)
	if u.Password == u2.Password {
		c.SetSession("uid", u2.Uuid)
		logs.Info("login success")
		c.Redirect("/", 302)
	} else {
		c.Get()
		c.Data["error"] = "Username / password does not match"
	}
}

func (c *LoginController) Get() {
	c.Data["Title"] = "Login"
	c.Layout = "layout.html"
	c.TplName = "login.html"
	c.Data["login"] = true
}

func (c *RegController) Post() {
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
		logs.Error(err)
		c.Redirect("/register", 302)
		c.Data["error"] = err.Error()
		return
	}

	if u.Username == "" || u.Password == "" {
		c.Get()
		c.Data["error"] = "Username / password should not be empty"
		return
	}

	u2 := models.GetUserByUserName(u.Username)
	if u2.Uuid != "" {
		c.Get()
		c.Data["error"] = "Username exists"
		return
	}

	u.Uuid = uuid.New().String()
	models.SaveUser(&u)
	c.SetSession("uid", u.Uuid)
	c.Redirect("/", 302)
}

func (c *RegController) Get() {
	c.Data["Title"] = "Register"
	c.Layout = "layout.html"
	c.TplName = "login.html"
}
