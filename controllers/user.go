package controllers

import (
	"fmt"
	"myorm/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

type HomeController struct {
	beego.Controller
}

// func (this *HomeController) Get() {
// 	this.Ctx.WriteString("Hellow beego")
// }
// 用户列表
func (this *HomeController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	userlist := models.DataList()
	fmt.Println(userlist)
	/*返回json数据*/
	this.Data["datas"]=userlist
	this.TplName = "home.html"
	
}
