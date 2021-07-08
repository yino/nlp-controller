package repository

import "nlp/domain/entity"

type UserRepository interface {
	Add(user *entity.User) error
	Edit(user *entity.User) error
	GetUserList(search map[string]interface{}) ([]entity.User, error)
	GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []entity.User, total uint, err error)
	UserInfo(uint64) (*entity.User, error)
	FindUserInfo(search map[string]interface{}) (*entity.User, error)
}

