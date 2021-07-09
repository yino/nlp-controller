package repository

import (
	"github.com/yino/nlp-controller/domain/po"
)

// user 聚合工厂
type UserRepository interface {
	Add(user *po.User) error
	Edit(user *po.User) error
	GetUserList(search map[string]interface{}) ([]po.User, error)
	GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []po.User, total uint, err error)
	UserInfo(uint64) (*po.User, error)
	FindUserInfo(search map[string]interface{}) (*po.User, error)
	FindUserByToken(token string)(*po.User, error)
}
