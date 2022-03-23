package db

import "awesomeProject1/model"

func InsertCourseSelectionInformation(selection *model.StudentCourseRelation) (*model.StudentCourseRelation, error) {
	err := Db.Create(selection).Error
	if err != nil {
		return nil, err
	}
	return selection, err
}
func GetCourseSelectionInformation(selections int) ([]model.Course, error) {
	relations := make([]model.StudentCourseRelation, 0)
	err := Db.Where("student_id=?", selections).Find(&relations).Error
	if err != nil {
		return nil, err
	}
	Req := make([]model.Course, 0)
	for _, relation := range relations {
		req := model.Course{}
		Err := Db.Where("course_id=?", relation.ClassID).Find(&req).Error
		Req = append(Req, req)
		if Err != nil {

		}
	}
	return Req, err
}
func CancelCourseSelectionInformation(studentID int, classID int) (*model.StudentCourseRelation, error) {
	var req *model.StudentCourseRelation = &model.StudentCourseRelation{}
	err := Db.Where("student_id=? AND class_id=?", studentID, classID).Delete(req).Error
	if err != nil {
		return nil, err
	}
	return req, err
}
func GetCourseSelectionNum(classID int) (int, error) {
	var count int
	err := Db.Model(&model.StudentCourseRelation{}).Where("class_id = ?", classID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
