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
	var loginReq loginReq
	if err := c.ShouldBindJSON(loginReq); err != nil {
		c.JSON(-1, "参数错误")
		return
	}
	search := make(map[string]interface{})
	search["mobile"] = loginReq.Mobile
	search["password"] = loginReq.Password
	result, err := u.us.FindUserInfo(search)
	if err != nil {
		c.JSON(-1, result)
	} else {
		c.JSON(200, result)
	}
	return
}

func (u *Users) HandlerUserRegister(c *gin.Context) {

}

func NewUsersInterface(us application.UserApp) Users {
	return Users{
		us: us,
	}
}
