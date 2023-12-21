package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	Id        int    `gorm:"primary_key,AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(250)"`
	Telephone string `gorm:"type:varchar(250)"`
	Address   string `gorm:"type:varchar(250)"`
}

func init() {
	configure := initConfigure()
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", configure.Mysql.Username, configure.Mysql.Password, fmt.Sprintf("%s:%s", configure.Mysql.URL, configure.Mysql.Port), configure.Mysql.Database))
	if err != nil {
		glog.Errorf("Connect to mysql failed: %s", err.Error())
	}
	db.SingularTable(true)
	db = db.AutoMigrate(&User{})
}
