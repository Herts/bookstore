package routers

import (
	"github.com/Herts/bookstore/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	//beego.ErrorController(&controllers.ErrorController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/book",
			beego.NSInclude(&controllers.BookController{}),
		),
	)
	beego.AddNamespace(ns)
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	//beego.InsertFilter("/*", beego.BeforeRouter, InitFilter())

	beego.Router("/", &controllers.MainController{})
	beego.Router("/book", &controllers.MainController{}, "get:ViewBookPage")

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegController{})

	beego.Router("/manage/editbook", &controllers.SellerController{}, "get:EditBookPage")
	beego.Router("/manage/addbook", &controllers.SellerController{}, "get:AddBookPage")
	beego.Router("/manage/books", &controllers.SellerController{}, "get:AllBooksPage")
	beego.Router("/manage/book", &controllers.SellerController{}, "get:ViewBookPage")

	beego.Router("/user/orders", &controllers.BuyerController{}, "get:AllOrdersPage")
	beego.Router("/user/order", &controllers.BuyerController{}, "get:OrderPage")
	beego.Router("/user/cart", &controllers.BuyerController{}, "get:CartPage")
	beego.Router("/user/addresses", &controllers.BuyerController{}, "get:AllAddressesPage")
	beego.Router("/user/editaddress", &controllers.BuyerController{}, "get:EditAddressPage")

}

func InitFilter() beego.FilterFunc {
	auth := beego.AppConfig.String("Authorization")
	return func(ctx *context.Context) {
		target := ctx.Input.URL()
		if strings.HasPrefix(target, "/login") || strings.HasPrefix(target, "/register") || target == "/" {
			return
		}
		if len(auth) != 0 {
			if ctx.Input.Header("Authorization") == auth {
				return
			}
		}
		_, ok := ctx.Input.Session("uid").(string)
		if !ok {
			ctx.Redirect(402, "/login")
		}
	}
}
