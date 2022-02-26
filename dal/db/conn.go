package db

import "github.com/jinzhu/gorm"

var Db *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", "root:yxqc20161012@/RUNOOB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	Db = db
}
