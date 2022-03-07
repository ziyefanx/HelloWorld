package courseSelection

type GetSelectionReq struct {
	StudentID int `form:"studentID" json:"studentID" uri:"studentID" binding:"required"`
	ClassID   int `form:"classID" json:"classID" uri:"classID"`
}
type DeleteSelectionReq struct {
	StudentID int `form:"studentID" json:"studentID" uri:"studentID" binding:"required"`
	ClassID   int `form:"classID" json:"classID" uri:"classID"`
}
type CreateSelectionReq struct {
	StudentID int `form:"studentID" json:"studentID" uri:"studentID" binding:"required"`
	ClassID   int `form:"classID" json:"classID" uri:"classID" binding:"required"`
}
type GetSelectionNumReq struct {
	ClassID int `form:"classID" json:"classID" uri:"classID" binding:"required"`
}
