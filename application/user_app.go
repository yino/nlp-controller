package application

import (
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/infrastructure/persistence"
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
		log.Error("user Add", err)
		return interfaces.ErrorRegister, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// Edit edit
func (u *UserApp) Edit(user *entity.User) error {
	err := u.userDomain.Edit(user)
	if err != nil {
		log.Error("user edit", err)
	}
	return err
}

// GetUserList get user list
func (u *UserApp) GetUserList(search map[string]interface{}) ([]vo.UserVo, error) {
	res, err := u.userDomain.GetUserList(search)
	if err != nil {
		log.Error("GetUserList", err)
	}
	return res, err
}

// GetUserPage get user list page
func (u *UserApp) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (pageVo vo.UserPageVo, err error) {
	pageVo, err = u.userDomain.GetUserPage(search, page, pageSize)
	if err != nil {
		log.Error("GetUserPage", err)
	}
	return
}

// UserInfo get user info by id
func (u *UserApp) UserInfo(id uint64) (vo.UserVo, int) {
	userVo, err := u.userDomain.UserInfo(id)
	if err != nil {
		log.Error("UserInfo", err)
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
func (u *UserApp) CreateAk(uid uint64, akType string) (int, string) {
	err := u.userDomain.CreateAppKey(uid, akType)
	if err != nil {
		log.Error("CreateAk", err)
		return interfaces.ErrorCreateData, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

// AkPage 获取ak page
func (u *UserApp) AkPage(uid uint64, createType string, page, pageSize uint) (ret int, vo vo.UserAkVoPage) {
	vo = u.userDomain.AppKeyPage(uid, createType, page, pageSize)
	return interfaces.StatusSuccess, vo
}

// DeleteUserAk delete ak
func (u *UserApp) DeleteUserAk(uid, id uint64) (int, string) {
	err := u.userDomain.DeleteAppKey(id, uid)
	if err != nil {
		log.Error("DeleteUserAk", err)
		return interfaces.ErrorDeleteData, err.Error()
	}
	return interfaces.StatusSuccess, ""
}

func (u *UserApp) FindAkByUser(ak string) (vo.UserVo, int) {
	user, err := u.userDomain.FindUserByAk(ak)
	if err != nil {
		log.Error("FindAkByUser", err)
		return user, interfaces.ErrorUserNotFound
	}
	return user, interfaces.StatusSuccess
}

// NewUserApp new user app
func NewUserApp(repo *persistence.Repositories) UserApp {
	return UserApp{
		userDomain: domain.NewUserDomain(repo.User),
	}
}
