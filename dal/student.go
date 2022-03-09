package db

import (
	"awesomeProject1/model"
)

func InsertStudentInformation(student *model.Student) (*model.Student, error) {
	err := Db.Create(student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}

func UpdateStudentInformation(student *model.Student) (*model.Student, error) {
	err := Db.Model(&model.Student{}).Updates(student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}
func DeleteStudentInformation(id uint) (*model.Student, error) {
	var student *model.Student = &model.Student{}
	Db.First(&student, id)
	err := Db.Delete(&student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}
func SelectStudentByID(id uint) (*model.Student, error) {
	var student = &model.Student{}
	err := Db.First(student, id).Error
	if err != nil {
		return nil, err
	}
	return student, err
}

func SelectStudentByName(name string) (*model.Student, error) {
	var student = &model.Student{}
	err := Db.Where("name LIKE ?", name).First(student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}
func SelectByGrade(grade int) (*model.Student, error) {
	var student = &model.Student{}
	err := Db.Where("grade=?", grade).First(student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}
func SelectBySex(sex uint) (*model.Student, error) {
	var student = &model.Student{}
	err := Db.Where("sex=?", sex).First(student).Error
	if err != nil {
		return nil, err
	}
	return student, err
}
