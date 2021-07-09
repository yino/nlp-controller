package vo

type PageVo struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
type UserLoginVo struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Mobile uint64 `json:"mobile"`
}

type UserVo struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Mobile uint64 `json:"mobile"`
	Email  string `json:"email"`
}

type UserPageVo struct {
	PageVo
	Data []UserVo
}

type QaQuestionVo struct {
	Id       uint64 `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type QaQuestionPageVo struct {
	PageVo
	Data []QaQuestionVo
}
type QaQuestionInfoVo struct {
	QaQuestionVo
	SimilarQuestion []QaQuestionVo
}
