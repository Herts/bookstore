package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB

func InitDb() {
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", viper.GetString("db.username"),
		viper.GetString("db.password"), viper.GetString("db.url"),
		viper.GetString("db.port"), viper.GetString("db.database"))
	var err error
	db, err = gorm.Open("mysql", dbLink)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(viper.GetBool("db.debug"))
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Address{})
}
