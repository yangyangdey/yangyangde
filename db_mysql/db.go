package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_"go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_pwd")
	dbIp := config.String("db_ip")
	dbName := config.String("db_Name")
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据连接错误")
	}
	DB = db

}
