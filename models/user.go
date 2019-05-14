package models

import (
	"fmt"
    "github.com/astaxie/beego/orm"
)

// 用户信息
type User struct {
	Id	int
	Name	string
	Pwd	string
	Email	string
	Sex	string
	Phone	string
}

//根据用户数据列表 
func DataList() (users []User) {
     
    o := orm.NewOrm()
    qs := o.QueryTable("user")
     
    var us []User
    cnt, err :=  qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
    if err == nil {
        fmt.Println("count", cnt)
    }
    return us
}

func init() {
    // 需要在init中注册定义的model
	// orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
	orm.RegisterModel(new(User))
}