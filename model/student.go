package model

// 存放数据模型

type StudentType string

const (
	StudentTypeNormal StudentType = "normal"
	StudentTypeArt    StudentType = "art"
)

type Student struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Sex         uint
	Grade       int
	StudentType StudentType
}
