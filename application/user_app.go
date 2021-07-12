package application

import (
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"
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
func (u *UserApp) UserInfo(id uint64) (vo.UserVo, int) {
	return u.userDomain.UserInfo(id)
}

func (u *UserApp) Login(search map[string]interface{}) (vo vo.UserLoginVo, ret int) {
	return u.userDomain.Login(search)
}

func (u *UserApp) AuthToken(token string) (vo.UserVo, int) {
	vo, ok := u.userDomain.AuthToken(token)
	if !ok {
		return vo, interfaces.ErrorToken
	} else {
		return vo, interfaces.StatusSuccess
	}
}

func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userDomain: domain.NewUserDomain(repo),
	}
}
