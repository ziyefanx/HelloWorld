package course

type GetCourseReq struct {
	CourseID     int    `form:"courseID" json:"courseID" uri:"courseID" binding:"required"`
	CourseName   string `form:"courseName" json:"courseName" uri:"courseName" binding:"required"`
	CourseCredit int    `form:"courseCredit" json:"courseCredit" uri:"courseCredit" binding:"required"`
}
