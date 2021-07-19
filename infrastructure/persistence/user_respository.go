package persistence

import (
	"errors"
	"strconv"

	"github.com/yino/common"
	"github.com/yino/nlp-controller/domain/po"

	"gorm.io/gorm"
)

// UserRepo user repo
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepository user test
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Add add
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

// Edit edit
func (obj *UserRepo) Edit(user *po.User) error {
	res := obj.db.Where("id = ?", user.ID).Save(user)
	return res.Error
}

// GetUserList get list
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

// GetUserPage get page
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

// UserInfo user info
func (obj *UserRepo) UserInfo(id uint64) (*po.User, error) {

	user := new(po.User)
	res := obj.db.First(user)

	return user, res.Error
}

// FindUserInfo find user info
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

// FindUserByToken find user info by token
func (obj *UserRepo) FindUserByToken(token string) (*po.User, error) {
	whereUser := new(po.User)
	user := new(po.User)
	user.Token = token
	res := obj.db.Where(whereUser).First(user)

	return user, res.Error
}

// CreateAk creat ak
func (obj *UserRepo) CreateAk(keyPo *po.UserAppKeyPo) error {
	return obj.db.Create(keyPo).Error
}

// GetAkPage get ak page
func (obj *UserRepo) GetAkPage(search map[string]interface{}, page, pageSize uint) (datList []po.UserAppKeyPo, total uint, err error) {
	var count int64
	db := obj.db
	if uid, ok := search["user_id"]; ok {
		db = db.Where("user_id = ?", uid)
	}
	err = db.Model(&po.UserAppKeyPo{}).Count(&count).Error
	if err != nil {
		return
	}
	if page > 0 {
		page = page - 1
	}
	err = db.Limit(int(pageSize)).Offset(int(page * pageSize)).Order("id desc").Find(&datList).Error
	if err != nil {
		return
	}
	total = uint(count)

	return
}

// FindUserAk find user ak
func (obj *UserRepo) FindUserAk(uid uint64, ak string, as string) (po.UserAppKeyPo, error) {

	return po.UserAppKeyPo{}, nil
}
