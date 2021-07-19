package domain

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"

	"github.com/yino/common"
	"go.uber.org/zap"
)

// User 用户领域服务
type User struct {
	UserRepo repository.UserRepository // user 聚合工厂
}

// Add add user
// @param user	entity *entity.User	"user 实体"
// @return int "是否成功"
// @return string "描述"
func (u *User) Add(user *entity.User) error {
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
		return err
	}
	return nil

}

// Edit edit user
// @param user	*entity.User	"user 实体"
// @return error
func (u *User) Edit(user *entity.User) error {
	userPo := new(po.User)
	userPo.Name = user.Name
	userPo.Mobile = user.Mobile
	userPo.Email = user.Email
	userPo.Password = user.Password
	userPo.ID = user.ID
	return u.UserRepo.Edit(userPo)
}

// GetUserList get user list
// @param search map[string]interface{} "搜索 例如：["mobile"=>"xxxxx","password"=>"xxx"]"
// @return []vo.UserVo
// @return error
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

// GetUserPage get user page
// @param search map[string]interface{} "搜索 例如：["mobile"=>"xxxxx","password"=>"xxx"]"
// @param page uint "start page"
// @param pageSize uint "limit"
// @return vo.UserPageVo
// @return error
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

// UserInfo get user info by id
// @param id uint64 "user id"
// @return vo.UserVo
// @return int
func (u *User) UserInfo(id uint64) (vo.UserVo, error) {
	userPo, err := u.UserRepo.UserInfo(id)
	if err != nil {
		return vo.UserVo{}, err
	}
	return vo.UserVo{
		Id:     userPo.ID,
		Mobile: userPo.Mobile,
		Name:   userPo.Name,
		Email:  userPo.Email,
	}, nil
}

// Login login
// @param search map[string]interface{} "搜索 例如：["mobile"=>"xxxxx","password"=>"xxx"]"
// return vo.UserLoginVo
// return ret
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
	} else {
		vo.Token = user.Token
		vo.Mobile = user.Mobile
		vo.Name = user.Name
	}

	return vo, interfaces.StatusSuccess
}

// AuthToken 校验token
func (u *User) AuthToken(token string) (vo vo.UserVo, ok bool) {
	user, err := u.UserRepo.FindUserByToken(token)
	if err != nil {
		return vo, false
	}

	// token 不存在
	if user.ID == 0 {
		return vo, false
	}

	// 过期
	if user.TokenExpire.Unix() < time.Now().Unix() {
		return vo, false
	}

	vo.Mobile = user.Mobile
	vo.Name = user.Name
	vo.Mobile = user.Mobile
	vo.Id = user.ID
	return vo, true
}

// CreateAppKey create app key
func (u *User) CreateAppKey(uid uint64, createType string) error {
	userAkPo := new(po.UserAppKeyPo)

	timeNow := time.Now().Unix()
	timeStrNow := strconv.FormatInt(timeNow, 10)
	uidStr := strconv.FormatInt(int64(timeNow), 10)
	userAkPo.Ak, _ = common.EncryptPassword(timeStrNow + uidStr)
	userAkPo.As, _ = common.EncryptPassword(timeStrNow + uidStr)
	userAkPo.ReqNum = 0
	userAkPo.Type = createType
	userAkPo.UserID = uid
	return u.UserRepo.CreateAk(userAkPo)
}

// AppKeyPage get app key page
func (u *User) AppKeyPage(uid uint64, createType string, page, pageSize uint) vo.UserAkVoPage {
	search := make(map[string]interface{})
	search["user_id"] = uid
	search["create_type"] = createType
	dataList, total, err := u.UserRepo.GetAkPage(search, page, pageSize)
	var resp vo.UserAkVoPage
	resp.Page = int64(page)
	resp.PageSize = int64(pageSize)

	if err != nil {
		return resp
	}
	resp.Total = int64(total)

	for _, val := range dataList {
		resp.Data = append(resp.Data, vo.UserAkVo{
			ID:     val.ID,
			Ak:     val.Ak,
			As:     val.As,
			Type:   val.Type,
			ReqNum: val.ReqNum,
		})
	}
	return resp
}

// AuthAppKey auth ak as
func (u *User) AuthAppKey(ak string, as string) error {
	akInfo, err := u.UserRepo.FindUserAkByAkAs(ak, as)
	if err != nil {
		return err
	}

	if akInfo.ID > 0 {
		return nil
	}
	return errors.New("auth ak fail")
}

// DeleteAppKey delete ak
func (u *User) DeleteAppKey(id uint64) error {
	return nil
}

// NewUserDomain new domain.User
// return domain.User
func NewUserDomain(repo repository.UserRepository) User {
	return User{
		UserRepo: repo,
	}
}
