package application

import (
	"nlp/domain/entity"
	"nlp/domain/repository"
	"nlp/interfaces"
)

type UserApp struct {
	userRepo repository.UserRepository
}

func (u *UserApp) Add(user *entity.User) (int, string) {
	err := u.userRepo.Add(user)
	if err != nil {
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
func (u *UserApp) Login(search map[string]interface{}) (*entity.User, int) {
	user, err := u.userRepo.FindUserInfo(search)
	if err != nil {
		return user, interfaces.ErrorUserNotFound
	}
	return user, interfaces.StatusSuccess
}

func NewUserApp(repo repository.UserRepository) UserApp {
	return UserApp{
		userRepo: repo,
	}
}
