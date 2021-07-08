package application

import (
	"encoding/json"
	"time"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/interfaces"

	"github.com/yino/common"
	"go.uber.org/zap"
)

type UserApp struct {
	userRepo repository.UserRepository
}

func (u *UserApp) Add(user *entity.User) (int, string) {
	err := u.userRepo.Add(user)
	if err != nil {
		data, _ := json.Marshal(user)
		log.Logger.Fatal("add user fail",
			zap.ByteString("data", data),
			zap.Error(err),
		)
		return interfaces.ErrorRegister, err.Error()
	} else {
		return interfaces.StatusSuccess, ""
	}

}
func (u *UserApp) Edit(user *entity.User) error {
	return u.userRepo.Edit(user)
}
func (u *UserApp) GetUserList(search map[string]interface{}) ([]entity.User, error) {
	return u.userRepo.GetUserList(search)
}
func (u *UserApp) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []entity.User, total uint, err error) {
	return u.userRepo.GetUserPage(search, page, pageSize)
}
func (u *UserApp) UserInfo(id uint64) (*entity.User, error) {
	return u.userRepo.UserInfo(id)
}

func (u *UserApp) Login(search map[string]interface{}) (user *entity.User, token string, ret int) {
	user, err := u.userRepo.FindUserInfo(search)
	if err != nil {
		return user, "", interfaces.ErrorUserNotFound
	}
	if user.ID ==  0{
		return nil,"",interfaces.ErrorUserNotFound
	}
	token = common.CreateUidToken(user.ID)
	tokenExpire := time.Now().Unix() + config.Conf.App.TokenExpire
	tokenTimeExpire := time.Unix(tokenExpire, 0)
	user.Token = token
	user.TokenExpire = &tokenTimeExpire
	err = u.userRepo.Edit(user)

	if err != nil {
		log.Logger.Fatal("update user token fail",
			zap.Uint64("user id", user.ID),
			zap.String("user token", token),
			zap.Time("user token expire", tokenTimeExpire),
			zap.Error(err),
		)
		return user, token, interfaces.ErrorCreateToken
	}

	return user, token, interfaces.StatusSuccess
}

func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userRepo: repo,
	}
}
