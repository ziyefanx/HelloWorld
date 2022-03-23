package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", "root:yxqc20161012@/RUNOOB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println("error：", err.Error())
		panic("连接数据库失败")
	}
	Db = db

}
