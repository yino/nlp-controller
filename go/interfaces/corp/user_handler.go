package corp

import (
	"github.com/gin-gonic/gin"
	"nlp/application"
	"nlp/domain/entity"
	"nlp/interfaces"
	"strconv"
	"time"
)

type Users struct {
	us application.UserApp
}

func (u *Users) HandlerUserInfo(c *gin.Context) {
}
func (u *Users) HandlerUserEdit(c *gin.Context) {

}
func (u *Users) HandlerUserLogin(c *gin.Context) {
	var loginReq LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	search := make(map[string]interface{})
	search["mobile"] = loginReq.Mobile
	search["password"] = loginReq.Password
	result, ret := u.us.Login(search)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret)
	} else {
		interfaces.SendResp(c, result, ret)
	}

	return
}

func (u *Users) HandlerUserRegister(c *gin.Context) {
	var registerReq RegisterReq
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	mobile, _ := strconv.Atoi(registerReq.Mobile)
	userEntity := new(entity.User)
	userEntity.Password = registerReq.Password
	userEntity.Mobile = uint64(mobile)
	userEntity.Email = registerReq.Email
	userEntity.Name = registerReq.Name
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()
	ret, errMsg := u.us.Add(userEntity)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, errMsg)
	} else {
		interfaces.SendResp(c, ret)
	}
	return
}

func NewUsersInterface(us application.UserApp) Users {
	return Users{
		us: us,
	}
}
