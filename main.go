package main

import (
	"DataCertProtict/db_mysql"
	_ "DataCertProtict/routers"
	"github.com/astaxie/beego"
)

func main() {
    db_mysql.Connect()
    beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()

}