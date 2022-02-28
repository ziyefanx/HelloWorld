package model

// 存放数据模型

type Student struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Sex   uint
	Grade int
}
