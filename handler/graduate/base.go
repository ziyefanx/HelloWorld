package graduate

type GetGraduateReq struct {
	StudentID int `form:"studentID" json:"studentID" uri:"studentID" binding:"required"`
}
