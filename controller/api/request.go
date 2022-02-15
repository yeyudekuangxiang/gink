package api

type GetUserInfoForm struct {
	ID int `json:"id" form:"id" binding:"gt=0" alias:"用户id"`
}
