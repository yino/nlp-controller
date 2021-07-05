package application

import (
	"nlp/domain/entity"
	"nlp/domain/repository"
)

type UserApp struct {
	userRepo repository.UserRepository
}

func (u *UserApp) Add(user *entity.User) error {
	return u.userRepo.Add(user)
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

func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userRepo: repo,
	}
}
