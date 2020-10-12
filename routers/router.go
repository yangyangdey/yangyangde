package routers

import (
	"DataCertProtict/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user_register", &controllers.RegisterControllers{})
	beego.Router("/longin.html", &controllers.LoginController{})
	beego.Router("/user_longin", &controllers.LoginController{})
}
