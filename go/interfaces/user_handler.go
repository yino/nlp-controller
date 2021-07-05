package interfaces

import (
	"github.com/gin-gonic/gin"
	"nlp/application"
)

type Users struct {
	us application.UserApp
}

func (u *Users) HandlerUserInfo(c *gin.Context) {

}
func (u *Users) HandlerUserEdit(c *gin.Context) {

}
func (u *Users) HandlerUserLogin(c *gin.Context) {

}

func (u *Users) HandlerUserRegister(c *gin.Context) {

}

func NewUsersInterface(us application.UserApp) Users {
	return Users{
		us:     us,
	}
}
