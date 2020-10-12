package models

import (
	"DataCertProtict/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int         `form:"id"`
	Phone string      `form:"phone"`
	Password string   `form:"password"`
}

func (u User)SaveUser() (int64,error) {
	md5Hash :=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordByte:=md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordByte)
	row,err:=db_mysql.DB.Exec("insert into beegouserfrom(phone,password) values(?,?)",u.Phone,u.Password)
	if err !=nil {
		return -1,err
	}
	id,err :=row.RowsAffected()
	if err !=nil {
		return -1,err
	}
	return id,nil

}
func (u User) QueryUser() (*User,error) {
	md5Hash :=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordByte:=md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordByte)
	row :=db_mysql.DB.QueryRow("select phone from beegouserfrom where phone= ? and password= ?",u.Phone, u.Password)
	 err:=row.Scan(&u.Phone)
	if err !=nil {
		return nil,err
	}
		return &u,nil
}