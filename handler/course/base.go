package course

type GetCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit"`
	CourseType        string `form:"courseType" json:"courseType" uri:"courseType"`
	Page              int    `form:"page" json:"page" uri:"page" binding:"required"`
	Size              int    `form:"size" json:"size" uri:"size" binding:"required"`
}
type UpdateCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName" binding:"required"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" binding:"required"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit" binding:"required"`
	CourseType        string `form:"courseType" json:"courseType" uri:"courseType" binding:"required"`
}
type DeleteCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit"`
	CourseType        string `form:"courseType" json:"courseType" uri:"courseType"`
}
type CreateCourseReq struct {
	CourseID          int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName        string `form:"courseName" json:"courseName" uri:"courseName" binding:"required"`
	CourseCredit      int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" binding:"required"`
	CourseNumberLimit int    `form:"courseNumberLimit" json:"courseNumberLimit" uri:"courseNumberLimit" binding:"required"`
	CourseType        string `form:"courseType" json:"courseType" uri:"courseType" binding:"required"`
}
