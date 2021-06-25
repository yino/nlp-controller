package repository

type LoginRepository interface {
	Login(username,  password string) (string, error)
	Register()
}
