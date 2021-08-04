package vo

import "time"

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
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	Mobile        uint64 `json:"mobile"`
	Email         string `json:"email"`
	QaModelStatus uint8  `json:"qa_model_status"`
}

type UserPageVo struct {
	PageVo
	Data []UserVo
}

type UserAkVo struct {
	ID        uint64    `json:"id"`
	Ak        string    `json:"app_key"`
	As        string    `json:"app_secret"`
	Type      string    `json:"type"`
	ReqNum    uint64    `json:"req_num"`
	CreatedAt time.Time `json:"created_at"`
}
type UserAkVoPage struct {
	PageVo
	Data []UserAkVo
}
type QaQuestionVo struct {
	Id       uint64 `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	UserId   uint64 `json:"user_id"`
}

type QaQuestionPageVo struct {
	PageVo
	Data []QaQuestionVo
}
type QaQuestionInfoVo struct {
	QaQuestionVo
	SimilarQuestion []QaQuestionVo
}
type QaMatchQuestionItemVo struct {
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	Sims     float64 `json:"sims"`
}
type QaMatchQuestionVo struct {
	Data []QaMatchQuestionItemVo
}
