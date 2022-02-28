package student

type GetStudentReq struct {
	ID    int    `form:"id" json:"id" uri:"id" xml:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" xml:"grade" binding:"required"`
}

type UpdateStudentInfoReq struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" xml:"grade" binding:"required"`
}
