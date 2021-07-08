package persistence

import (
	"errors"
	"gorm.io/gorm"
	"nlp/domain/entity"
	"strconv"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepository： user test
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (obj *UserRepo) Add(user *entity.User) error {
	var count int64
	err := obj.db.Model(&entity.User{}).Where("mobile = ?", user.Mobile).Count(&count).Error
	if err != nil {
		return err
	}
	if count >0  {
		return errors.New("mobile already exists")
	}
	res := obj.db.Create(user)
	return res.Error
}
func (obj *UserRepo) Edit(user *entity.User) error {
	res := obj.db.Save(user)
	return res.Error
}
func (obj *UserRepo) GetUserList(search map[string]interface{}) ([]entity.User, error) {
	var userList []entity.User
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

func (obj *UserRepo) GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []entity.User, total uint, err error) {
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
func (obj *UserRepo) UserInfo(id uint64) (*entity.User, error) {

	user := new(entity.User)
	res := obj.db.First(user)

	return user, res.Error
}

func (obj *UserRepo) FindUserInfo(search map[string]interface{}) (*entity.User, error) {
	user := new(entity.User)
	if searchMobile, ok := search["mobile"]; ok {
		intMobile, _ := strconv.Atoi(searchMobile.(string))
		mobile := uint64(intMobile)
		user.Mobile = mobile
	}

	if password, ok := search["password"]; ok {
		user.Password = password.(string)
	}
	res := obj.db.First(user)
	return user, res.Error
}

