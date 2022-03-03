package db

import (
	"awesomeProject1/model"
)

func InsertCourseInformation(course *model.Course) (*model.Course, error) {
	err := Db.Create(course).Error
	if err != nil {
		return nil, err
	}
	return course, err
}

func UpdateCourseInformation(course *model.Course) (*model.Course, error) {
	err := Db.Model(&model.Course{}).Updates(course).Error
	if err != nil {
		return nil, err
	}
	return course, err
}
func DeleteCourseInformation(id uint) (*model.Course, error) {
	var course *model.Course = &model.Course{}
	Db.First(&course, id)
	err := Db.Delete(&course).Error
	if err != nil {
		return nil, err
	}
	return course, err
}

func SelectCourseByID(id uint) (*model.Course, error) {
	var course model.Course
	err := Db.Where("course=?", id).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, err
}

func SelectCourseByName(name string) (*model.Course, error) {
	var course model.Course
	err := Db.Where("course_name=?", name).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, err
}
func SelectByCredit(courseCredit int) (*model.Course, error) {
	var course model.Course
	err := Db.Where("course_credit=?", courseCredit).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, err
}
