package student

// 存放公共结构体

// Login 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Opt   string `form:"opt" json:"opt" uri:"opt" xml:"opt" binding:"required"`
	ID    int    `form:"id" json:"id" uri:"id" xml:"id" binding:"required"`
	Name  string `form:"name" json:"name" uri:"name" xml:"name" binding:"required"`
	Sex   uint   `form:"sex" json:"sex" uri:"sex" xml:"sex" binding:"required"`
	Grade int    `form:"grade" json:"grade" uri:"grade" xml:"grade" binding:"required"`
}

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
