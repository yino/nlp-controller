package vo

type LoginReq struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type RegisterReq struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
	Name     string `form:"name" json:"name" binding:"required"`
}

type EditUserReq struct {
	RegisterReq
	Id uint64 `form:"id" json:"id" binding:"required"`
}
