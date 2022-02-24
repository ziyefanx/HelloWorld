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

func init() {
	db, err := gorm.Open("mysql", "root:yxqc20161012@/RUNOOB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	Db = db
}
func InsertStudentInformation() {
	var name string
	var sex uint
	var grade int
	fmt.Println("请输入要插入学生的信息：名字，性别，成绩")
	fmt.Scan(&name, &sex, &grade)
	Db.Create(&Student{
		Name:  name,
		Sex:   sex,
		Grade: grade,
	})
	fmt.Println("插入成功:")
}
func UpdateStudentInformation() {
	var name string
	var id, sex uint
	var grade int
	fmt.Println("请输入要修改学生的信息：学号，修改后的名字,性别，成绩")
	fmt.Scan(&id, &name, &sex, &grade)
	var student Student
	Db.First(&student, id)
	if student.ID == 0 {
		fmt.Println("查无此人，修改失败")
		return
	}
	Db.Model(&student).Updates(Student{Name: name, Sex: sex, Grade: grade})
	fmt.Println("更新完毕")
}
func DeleteStudentInformation() {
	var id uint
	fmt.Println("请输入要删除学生的学号")
	fmt.Scan(&id)
	var student Student
	Db.First(&student, id)
	if student.ID == 0 {
		fmt.Println("查无此人，删除失败")
		return
	}
	Db.Delete(&student)
	fmt.Println("删除成功")
}
func SelectStudentInformation() {
	fmt.Println("请输入查询的方式：" +
		"1.按学号查询" +
		"2.按姓名查询" +
		"3.按成绩查询")
	var choose int
	fmt.Scan(&choose)
	switch choose {
	case 1:
		SelectByID()
	case 2:
		SelectByName()
	case 3:
		SelectByGrade()
	default:
		fmt.Println("请重新输入")
	}
}
func SelectByID() {
	var id uint
	fmt.Println("请输入学生学号：")
	fmt.Scan(&id)
	var student Student
	Db.First(&student, id)
	if student.ID == 0 {
		fmt.Println("查无此人")
		return
	}
	fmt.Println(student)
}
func SelectByName() {
	var name string
	fmt.Println("请输入学生姓名：")
	fmt.Scan(&name)
	var student Student
	Db.Where("name=?", name).First(&student)
	if student.ID == 0 {
		fmt.Println("查无此人")
		return
	}
	fmt.Println(student)
}
func SelectByGrade() {
	var grade int
	fmt.Println("请输入学生成绩：")
	fmt.Scan(&grade)
	var student Student
	Db.Where("grade=?", grade).First(&student)
	if student.ID == 0 {
		fmt.Println("查无此人")
		return
	}
	fmt.Println(student)
}
func main() {
	Db.AutoMigrate(&Student{})
loop:
	for true {
		var choose int
		fmt.Println("请输入你的选择：" +
			"1.增加学生信息" +
			"2.删除学生信息" +
			"3.修改学生信息" +
			"4.查询学生信息" +
			"5.退出管理系统")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			InsertStudentInformation()
		case 2:
			DeleteStudentInformation()
		case 3:
			UpdateStudentInformation()
		case 4:
			SelectStudentInformation()
		case 5:
			break loop
		default:
			fmt.Println("请重新输入")
		}
	}
}
