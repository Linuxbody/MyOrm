package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ 下划线，只是执行一下这个包的init函数
	_ "github.com/go-sql-driver/mysql"
    _ "MyOrm/routers"
)

func init() {
	mysqlhost := beego.AppConfig.String("mysqlhost")
    mysqlport := beego.AppConfig.String("mysqlport")
    mysqluser := beego.AppConfig.String("mysqluser")
    mysqlpass := beego.AppConfig.String("mysqlpass")
    mysqldb := beego.AppConfig.String("mysqldb")
    
    //注册mysql Driver
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //注册数据库直接连接
    // orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
    
     //用户名:密码@tcp(url地址)/数据库 ,名字对应app.conf配置数据库信息
    conn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlhost + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8"
    //注册数据库连接
    orm.RegisterDataBase("default", "mysql", conn)
 
    fmt.Printf("数据库连接成功！%s\n", conn)
    // 自动建表
    // default,使用默认数据库
    // 默认false, 是否更新数据库表结构
    // 默认true，是否显示创建过程
    orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

