package persistence

import (
	"errors"
	"strconv"

	"github.com/yino/common"
	"github.com/yino/nlp-controller/domain/po"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepository： user test
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (obj *UserRepo) Add(user *po.User) error {
	var count int64
	err := obj.db.Model(&po.User{}).Where("mobile = ?", user.Mobile).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("mobile already exists")
	}
	res := obj.db.Create(user)
	return res.Error
}
func (obj *UserRepo) Edit(user *po.User) error {
	res := obj.db.Where("id = ?", user.ID).Save(user)
	return res.Error
}
func (obj *UserRepo) GetUserList(search map[string]interface{}) ([]po.User, error) {
	var userList []po.User
	db := obj.db
	if name, ok := search["name"]; ok {
		db = db.Where("name like %?%", name)
	}
	err := db.Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (obj *UserRepo) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []po.User, total uint, err error) {
	var count int64
	db := obj.db
	if name, ok := search["name"]; ok {
		db = db.Where("name like %?%", name)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}

	err = db.Find(&datList).Error
	if err != nil {
		return
	}
	total = uint(count)

	return
}
func (obj *UserRepo) UserInfo(id uint64) (*po.User, error) {

	user := new(po.User)
	res := obj.db.First(user)

	return user, res.Error
}

func (obj *UserRepo) FindUserInfo(search map[string]interface{}) (*po.User, error) {
	whereUser := new(po.User)
	user := new(po.User)
	if searchMobile, ok := search["mobile"]; ok {
		intMobile, _ := strconv.Atoi(searchMobile.(string))
		mobile := uint64(intMobile)
		whereUser.Mobile = mobile
	}

	if password, ok := search["password"]; ok {
		whereUser.Password = common.MD5(password.(string))
	}
	res := obj.db.Where(whereUser).First(user)

	return user, res.Error
}

func (obj *UserRepo) FindUserByToken(token string) (*po.User, error) {
	whereUser := new(po.User)
	user := new(po.User)
	user.Token = token
	res := obj.db.Where(whereUser).First(user)

	return user, res.Error
}
