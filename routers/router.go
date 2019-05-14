package routers

import (
	"MyOrm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{})
}
