package interfaces

type loginReq struct {
	Mobile   string `form:"user" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
