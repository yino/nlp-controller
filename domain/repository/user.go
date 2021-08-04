package repository

import (
	"github.com/yino/nlp-controller/domain/po"
)

// UserRepository 聚合工厂
type UserRepository interface {
	Add(user *po.User) error
	Edit(user *po.User) error
	GetUserList(search map[string]interface{}) ([]po.User, error)
	GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []po.User, total uint, err error)
	UserInfo(uid uint64) (*po.User, error)
	FindUserInfo(search map[string]interface{}) (*po.User, error)
	FindUserByToken(token string) (*po.User, error)
	CreateAk(keyPo *po.UserAppKeyPo) error
	GetAkPage(search map[string]interface{}, page, pageSize uint) (datList []po.UserAppKeyPo, total uint, err error)
	FindUserAkByAkAs(ak string, as string) (po.UserAppKeyPo, error)
	FindUserAkByID(ID uint64) (*po.UserAppKeyPo, error)
	DeleteAkByID(ID uint64) error
	FindAkByUidType(uid uint64, createType string) (po.UserAppKeyPo, error)
	UpdateUserQaModel(uid uint64, ok bool) error
}
