package vo

// LoginReq login
type LoginReq struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// RegisterReq register
type RegisterReq struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
	Name     string `form:"name" json:"name" binding:"required"`
}

// EditUserReq edit user
type EditUserReq struct {
	RegisterReq
	ID uint64 `form:"id" json:"id" binding:"required"`
}

// QaAddReq add qa
type QaAddReq struct {
	Question      string   `json:"question" binding:"required"`
	Answer        string   `json:"answer" binding:"required"`
	SlaveQuestion []string `json:"slave_question" bind:"required"`
}

type slaveQuestion struct {
	ID       uint64 `form:"id" json:"id" binding:"required"`
	Question string `json:"question" binding:"required"`
}

// QaEditReq edit qa
type QaEditReq struct {
	ID            uint64          `form:"id" json:"id" binding:"required"`
	Question      string          `json:"question" binding:"required"`
	Answer        string          `json:"answer" binding:"required"`
	SlaveQuestion []slaveQuestion `json:"slave_question" bind:"required"`
}
