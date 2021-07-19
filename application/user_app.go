package application

import (
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"
)

// UserApp 用户领域服务
type UserApp struct {
	userDomain domain.User
}

// Add add
func (u *UserApp) Add(user *entity.User) (int, string) {
	err := u.userDomain.Add(user)
	if err != nil {
		return interfaces.ErrorRegister, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// Edit edit
func (u *UserApp) Edit(user *entity.User) error {
	return u.userDomain.Edit(user)
}

// GetUserList get user list
func (u *UserApp) GetUserList(search map[string]interface{}) ([]vo.UserVo, error) {
	return u.userDomain.GetUserList(search)
}

// GetUserPage get user list page
func (u *UserApp) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (pageVo vo.UserPageVo, err error) {
	return u.userDomain.GetUserPage(search, page, pageSize)
}

// UserInfo get user info by id
func (u *UserApp) UserInfo(id uint64) (vo.UserVo, int) {
	userVo, err := u.userDomain.UserInfo(id)
	if err != nil {
		return userVo, interfaces.ErrorUserNotFound
	}
	return userVo, interfaces.StatusSuccess
}

// Login login
func (u *UserApp) Login(search map[string]interface{}) (vo vo.UserLoginVo, ret int) {
	return u.userDomain.Login(search)
}

// AuthToken auth token.
func (u *UserApp) AuthToken(token string) (vo.UserVo, int) {
	vo, ok := u.userDomain.AuthToken(token)
	if !ok {
		return vo, interfaces.ErrorToken
	}
	return vo, interfaces.StatusSuccess
}

// CreateAk 创建 ak
func (u *UserApp) CreateAk(uid uint64, akType string) error {
	u.userDomain.CreateAppKey()
}

// AkPage 获取ak page
func (u *UserApp) AkPage() (vo vo.UserAkVoPage, ret int) {

}

// NewUserApp new user app
func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userDomain: domain.NewUserDomain(repo),
	}
}
