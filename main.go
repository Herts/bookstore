package main

import (
	"github.com/Herts/bookstore/models"
	_ "github.com/Herts/bookstore/routers"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	models.InitDb()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
