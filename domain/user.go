package domain

import (
	"encoding/json"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/vo"
	"time"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/interfaces"

	"github.com/yino/common"
	"go.uber.org/zap"
)

// 用户领域服务
type User struct {
	UserRepo repository.UserRepository
}

func (u *User) Add(user *entity.User) (int, string) {
	userPo := new(po.User)
	userPo.Name = user.Name
	userPo.Mobile = user.Mobile
	userPo.Email = user.Email
	userPo.Password = user.Password
	err := u.UserRepo.Add(userPo)
	if err != nil {
		data, _ := json.Marshal(userPo)
		log.Logger.Error("add user fail",
			zap.ByteString("data", data),
			zap.Error(err),
		)
		return interfaces.ErrorRegister, err.Error()
	} else {
		return interfaces.StatusSuccess, ""
	}
}

func (u *User) Edit(user *entity.User) error {
	userPo := new(po.User)
	userPo.Name = user.Name
	userPo.Mobile = user.Mobile
	userPo.Email = user.Email
	userPo.Password = user.Password
	userPo.ID = user.ID
	return u.UserRepo.Edit(userPo)
}

func (u *User) GetUserList(search map[string]interface{}) ([]vo.UserVo, error) {

	list, err := u.UserRepo.GetUserList(search)
	if err != nil || len(list) == 0 {
		return []vo.UserVo{}, err
	}

	var userList []vo.UserVo
	for _, userPo := range list {
		userList = append(userList, vo.UserVo{
			Id:     userPo.ID,
			Mobile: userPo.Mobile,
			Name:   userPo.Name,
			Email:  userPo.Email,
		})
	}

	return userList, nil
}

func (u *User) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (vo.UserPageVo, error) {
	list, total, err := u.UserRepo.GetUserPage(search, page, pageSize)
	var userPageVo vo.UserPageVo
	if err != nil || total == 0 {
		return userPageVo, err
	}

	var dataList []vo.UserVo
	for _, userPo := range list {
		dataList = append(dataList, vo.UserVo{
			Id:     userPo.ID,
			Mobile: userPo.Mobile,
			Name:   userPo.Name,
			Email:  userPo.Email,
		})
	}
	userPageVo.Data = dataList
	userPageVo.Page = int64(page)
	userPageVo.Total = int64(total)
	userPageVo.PageSize = int64(pageSize)
	return userPageVo, err
}

func (u *User) UserInfo(id uint64) (vo.UserVo, error) {
	userPo, err := u.UserRepo.UserInfo(id)
	return vo.UserVo{
		Id:     userPo.ID,
		Mobile: userPo.Mobile,
		Name:   userPo.Name,
		Email:  userPo.Email,
	}, err
}

func (u *User) Login(search map[string]interface{}) (vo vo.UserLoginVo, ret int) {
	user, err := u.UserRepo.FindUserInfo(search)
	var token string
	if err != nil {
		return vo, interfaces.ErrorUserNotFound
	}
	if user.ID == 0 {
		return vo, interfaces.ErrorUserNotFound
	}
	// 校验token 是否存在及是否过期 不存在则生成新的token
	if user.TokenExpire.Unix() <= time.Now().Unix() {
		token = common.CreateUidToken(user.ID)

		tokenExpire := time.Now().Unix() + config.Conf.App.TokenExpire
		tokenTimeExpire := time.Unix(tokenExpire, 0)
		user.Token = token
		user.TokenExpire = &tokenTimeExpire

		err = u.UserRepo.Edit(user)
		if err != nil {
			log.Logger.Error("update user token fail",
				zap.Uint64("user id", user.ID),
				zap.String("user token", token),
				zap.Time("user token expire", tokenTimeExpire),
				zap.Error(err),
			)
			return vo, interfaces.ErrorCreateToken
		}
		vo.Token = token
		vo.Mobile = user.Mobile
		vo.Name = user.Name
	}

	return vo, interfaces.StatusSuccess
}

// AuthToken: 校验token
func (u *User) AuthToken(token string) {
	search := make(map[string]interface{})
	search["token"] = token
	user, err := u.UserRepo.FindUserInfo(search)
}

func NewUserDomain(repo repository.UserRepository) User {
	return User{
		UserRepo: repo,
	}
}
