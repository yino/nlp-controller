package persistence

import (
	"errors"
	"fmt"
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
	return obj.db.Where("id = ?", user.ID).Save(user).Error
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
	res := obj.db.Where("id = ?", id).First(user)

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
		if mobile == 0 {
			return user, errors.New("手机号格式错误")
		}
	}

	if searchName, ok := search["name"]; ok {
		whereUser.Name = searchName.(string)
		if len(whereUser.Name) == 0 {
			return user, errors.New("用户名错误")
		}
	}

	if password, ok := search["password"]; ok {
		whereUser.Password = common.MD5(password.(string))
	}
	res := obj.db.Debug().Where(whereUser).First(user)
	fmt.Println(user, res.Error)
	return user, res.Error
}

// FindUserByToken find user FindUserInfo by token
func (obj *UserRepo) FindUserByToken(token string) (*po.User, error) {
	user := new(po.User)
	res := obj.db.Where("token = ?", token).First(user)
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
	countDb := db.Model(&po.UserAppKeyPo{})
	if uid, ok := search["user_id"]; ok {
		db = db.Where("user_id = ?", uid)
		countDb = countDb.Where("user_id = ?", uid)
	}
	err = countDb.Count(&count).Error
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

// FindUserAkByID get ak info by id
func (obj *UserRepo) FindUserAkByID(id uint64) (*po.UserAppKeyPo, error) {
	userAk := new(po.UserAppKeyPo)
	err := obj.db.Where("id = ?", id).First(&userAk).Error
	return userAk, err
}

// FindUserAk find user ak
func (obj *UserRepo) FindUserAkByAkAs(ak string, as string) (po.UserAppKeyPo, error) {
	var userAk po.UserAppKeyPo
	err := obj.db.Where("app_key = ?", ak).Where("app_secret = ?", as).First(&userAk).Error
	return userAk, err
}

// FindAkByUidType find user ak by uid and type
func (obj *UserRepo) FindAkByUidType(uid uint64, createType string) (po.UserAppKeyPo, error) {
	var data po.UserAppKeyPo
	err := obj.db.Where("user_id = ?", uid).Where("type = ?", createType).First(&data).Error
	return data, err
}
func (obj *UserRepo) DeleteAkByID(id uint64) error {
	return obj.db.Where("id = ?", id).Delete(&po.UserAppKeyPo{}).Error
}

func (obj *UserRepo) UpdateUserQaModel(id uint64, ok bool) error {
	var status int8
	if ok {
		status = 1
	}
	return obj.db.Model(&po.User{}).Where("id", id).Update("qa_model_status", status).Error
}

func (obj *UserRepo) FindUserByAk(ak string) (*po.User, error) {
	var userAk po.UserAppKeyPo
	user := new(po.User)
	err := obj.db.Where("app_key = ?", ak).First(&userAk).Error

	if err != nil {
		return user, err
	}
	err = obj.db.Where("id = ?", userAk.UserID).First(user).Error
	return user, err
}
