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

func SelectCourseByLimit(limit uint, page int, size int) ([]model.Course, error) {
	courses := make([]model.Course, 0)
	pageindex := page
	pagesize := size
	err := Db.Offset((pageindex-1)*pagesize).Limit(pagesize).Where("course_number_limit=?", limit).Find(&courses).Error
	//err := Db.Where("course_number_limit=?", limit).Find(&course).Error
	if err != nil {
		return nil, err
	}
	return courses, err
}

func SelectCourseByName(name string, page int, size int) ([]model.Course, error) {
	courses := make([]model.Course, 0)
	pageindex := page
	pagesize := size
	Name := "%" + name + "%"
	err := Db.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&courses, "course_name LIKE ?", Name).Error
	//err := Db.Where("course_name LIKE ?", name).Find(&course).Error
	if err != nil {
		return nil, err
	}
	return courses, err
}
func SelectByCredit(courseCredit int, page int, size int) ([]model.Course, error) {
	courses := make([]model.Course, 0)
	pageindex := page
	pagesize := size
	err := Db.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&courses, "course_credit=?", courseCredit).Error
	//err := Db.Where("course_credit=?", courseCredit).Find(&course).Error
	if err != nil {
		return nil, err
	}
	return courses, err
}
func GetCourseLimit(id int) (int, error) {
	var course model.Course
	err := Db.First(&course, id).Error
	if err != nil {
		return 0, err
	}
	return course.CourseNumberLimit, err
}
