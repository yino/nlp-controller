package entity

import (
	"github.com/yino/nlp-controller/domain/po"
)

// User entity user
type User struct {
	po.User
	AppKey      []po.UserAppKeyPo `gorm:"foreignKey:UserId;"`
	QaQuestions []po.QaQuestion   `gorm:"foreignKey:UserId;"`
}

var AkType = []string{"QA"}
