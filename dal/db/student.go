package db

import (
	"awesomeProject1/model"
	"fmt"
)

func InsertStudentInformation(student *model.Student) (*model.Student, error) {
	err := Db.Create(student).Error
	return student, err
}

func UpdateStudentInformation(student *model.Student) (*model.Student, error) {
	err := Db.Model(&model.Student{}).Updates(student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
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
func SelectStudentByID(id int) (*model.Student, error) {
	var student *model.Student
	err := Db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}
func SelectStudentByName(name string) {
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
