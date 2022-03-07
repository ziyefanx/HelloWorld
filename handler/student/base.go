package student

type GetStudentReq struct {
	ID    uint   `form:"id" json:"id" uri:"id"`
	Name  string `form:"name" json:"name" uri:"name"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex"`
	Grade int    `form:"grade" json:"grade" uri:"grade"`
}

type UpdateStudentInfoReq struct {
	ID    uint   `form:"id" json:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name"  binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex"  binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" binding:"required"`
}
type DeleteStudentInfoReq struct {
	ID    uint   `form:"id" json:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex"`
	Grade int    `form:"grade" json:"grade" uri:"grade"`
}
type CreateStudentInfoReq struct {
	ID    uint   `form:"id" json:"id" uri:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name"  binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex"  binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" binding:"required"`
}
