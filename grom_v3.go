package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Sex   uint
	Grade int
}

var Db *gorm.DB

var stu Student

func Init() {
	db, err := gorm.Open("mysql", "root:yxqc20161012@/RUNOOB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	Db = db
}

func InsertStudentInformation(name string, sex uint, grade int) {
	Db.Create(&Student{
		Name:  name,
		Sex:   sex,
		Grade: grade,
	})
}

func UpdateStudentInformation(name string, id, grade int, sex uint) {
	var student Student
	Db.First(&student, id)
	Db.Model(&student).Updates(Student{Name: name, Sex: sex, Grade: grade})
}
func DeleteStudentInformation(id int) {
	var student Student
	Db.First(&student, id)
	err := Db.Delete(&student)
	if err != nil {
		fmt.Println("出现错误", err)
		return
	}
}
func SelectByID(id int) {
	var student Student
	Db.First(&student, id)
	stu = student
}
func SelectByName(name string) {
	var student Student
	Db.Where("name=?", name).First(&student)
	stu = student
}
func SelectByGrade(grade int) {
	var student Student
	Db.Where("grade=?", grade).First(&student)
	stu = student
}
func SelectBySex(sex uint) {
	var student Student
	Db.Where("sex=?", sex).First(&student)
	stu = student
}
