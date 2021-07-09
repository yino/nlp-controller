package application

import (
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
)

// 用户领域服务
type UserApp struct {
	userDomain domain.User
}

func (u *UserApp) Add(user *entity.User) (int, string) {
	return u.userDomain.Add(user)
}
func (u *UserApp) Edit(user *entity.User) error {
	return u.userDomain.Edit(user)
}
func (u *UserApp) GetUserList(search map[string]interface{}) ([]vo.UserVo, error) {
	return u.userDomain.GetUserList(search)
}
func (u *UserApp) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (pageVo vo.UserPageVo, err error) {
	return u.userDomain.GetUserPage(search, page, pageSize)
}
func (u *UserApp) UserInfo(id uint64) (vo.UserVo, error) {
	return u.userDomain.UserInfo(id)
}

func (u *UserApp) Login(search map[string]interface{}) (vo vo.UserLoginVo, ret int) {
	return u.userDomain.Login(search)
}

func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userDomain: domain.NewUserDomain(repo),
	}
}
