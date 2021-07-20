package corp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/interfaces"

	"github.com/gin-gonic/gin"
)

// Users user handler
type Users struct {
	us application.UserApp
}

// HandlerUserInfo get user info handler
// @Summary  获取用户信息
// @Description 获取用户信息
// @Tags corp
// @accept  json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security middleware.AuthToken
// @Success 200 {object} vo.UserVo
// @Router /v1/core/user/info [post]
func (u *Users) HandlerUserInfo(c *gin.Context) {
	id := c.GetUint64("uid")
	if id == 0 {
		interfaces.SendResp(c, interfaces.ErrorToken)
		return
	}
	vo, ret := u.us.UserInfo(id)
	interfaces.SendResp(c, vo, ret)
}

// HandlerUserEdit edit user
// @Summary  编辑用户
// @Description 编辑用户信息
// @Tags corp
// @accept  json
// @Produce json
// @Param login body vo.EditUserReq true "login"
// @Success 200
// @Router /v1/core/user/edit [post]
func (u *Users) HandlerUserEdit(c *gin.Context) {

}

// HandlerUserLogin login
// @Summary  corp登录
// @Description corp登录
// @Tags corp
// @accept  json
// @Produce json
// @Param login body vo.LoginReq true "login"
// @Success 200 {object} vo.UserLoginResp
// @Router /v1/core/login [post]
func (u *Users) HandlerUserLogin(c *gin.Context) {
	var loginReq vo.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	search := make(map[string]interface{})
	search["mobile"] = loginReq.Mobile
	search["password"] = loginReq.Password
	vo, ret := u.us.Login(search)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret)
		return
	}
	interfaces.SendResp(c, vo, ret)
}

// HandlerUserRegister register
// @Summary  corp注册
// @Description corp注册
// @Tags corp
// @accept  json
// @Produce json
// @Param login body vo.RegisterReq true "register"
// @Success 200
// @Router /v1/core/register [post]
func (u *Users) HandlerUserRegister(c *gin.Context) {
	var registerReq vo.RegisterReq
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
		return
	}
	interfaces.SendResp(c, ret)
}

// HandlerUserCreateAk 创建 ak
func (u *Users) HandlerUserCreateAk(c *gin.Context) {
	var createAkReq vo.CreateUserAkReq
	if err := c.ShouldBindJSON(&createAkReq); err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, err.Error())
		return
	}
	modelType := createAkReq.Type
	fmt.Println("modelType", modelType)
	uid := GetUid(c)
	ret, errMsg := u.us.CreateAk(uid, modelType)
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, errMsg)
		return
	}
	interfaces.SendResp(c, ret)
}

// HandlerUserAkPage ak page
func (u *Users) HandlerUserAkPage(c *gin.Context) {
	uid := GetUid(c)
	modelType := c.Query("type")
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams)
	}
	ret, data := u.us.AkPage(uid, modelType, uint(page), uint(limit))
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret)
		return
	}
	interfaces.SendResp(c, ret, data)
}

// HandlerUserAkDelete ak page
func (u *Users) HandlerUserAkDelete(c *gin.Context) {
	uid := GetUid(c)
	queryID := c.Query("id")
	id, err := strconv.ParseInt(queryID, 10, 64)
	if err != nil {
		interfaces.SendResp(c, interfaces.ErrorParams, "id fail")
		return
	}
	ret, errMsg := u.us.DeleteUserAk(uid, uint64(id))
	if ret != interfaces.StatusSuccess {
		interfaces.SendResp(c, ret, errMsg)
		return
	}
	interfaces.SendResp(c, ret)
}

// NewUsersInterface new UserInterface
func NewUsersInterface(us application.UserApp) Users {
	return Users{
		us: us,
	}
}
