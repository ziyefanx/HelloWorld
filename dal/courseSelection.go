package db

import "awesomeProject1/model"

func InsertCourseSelectionInformation(selection *model.StudentCourseRelation) (*model.StudentCourseRelation, error) {
	err := Db.Create(selection).Error
	if err != nil {
		return nil, err
	}
	return selection, err
}
func GetCourseSelectionInformation(selections int) (*model.StudentCourseRelation, error) {
	var req model.StudentCourseRelation
	err := Db.Where("student_id=?", selections).Find(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, err
}
func CancelCourseSelectionInfomation(classID int) (*model.StudentCourseRelation, error) {
	var req model.StudentCourseRelation
	Db.Where("class_id=?", classID).First(&req)
	err := Db.Delete(&req).Error
	if err != nil {
		return nil, err
	}
	return &req, err
}