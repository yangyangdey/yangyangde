package controllers

import (
	"DataCertProtict/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
func (l*LoginController) Get(){
	l.TplName = "longin.html"
}
func (l* LoginController) Post(){
  var user models.User
	err :=l.ParseForm(&user)
	fmt.Println(user.Phone,user.Password)
	if err !=nil{
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试！")
		return
	}
	u,err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	l.Data["phone"] =u.Phone
	l.TplName="home.html"
}