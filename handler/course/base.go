package course

type GetCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit"`
}
type UpdateCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName" binding:"required"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" binding:"required"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit" binding:"required"`
}
type DeleteCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit"`
}
type CreateCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName" binding:"required"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" binding:"required"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit" binding:"required"`
}
