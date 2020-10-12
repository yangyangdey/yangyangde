package controllers

import (
	"DataCertProtict/models"
	"github.com/astaxie/beego"
)

type RegisterControllers struct {
	beego.Controller
}

func (r *RegisterControllers) Post() {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，解析数据错误，请重试！")
		return
	}
	_, err = user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}
	r.TplName = "longin.html"
}
