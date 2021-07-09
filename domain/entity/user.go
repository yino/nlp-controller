package entity

import (
	"github.com/yino/nlp-controller/domain/po"
)

type User struct {
	po.User
	AppKey []*po.UserAppKeyPo
}


// CreateAppKey: 生成 app key
// 实体创建app key
func (user *User) CreateAppKey(){

}
