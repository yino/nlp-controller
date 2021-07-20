package domain_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/infrastructure/persistence"
)

var user domain.User

func init() {
	os.Setenv("env", "test")
	fmt.Println(os.Getenv("env"))
	config.GetConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	log.InitLogger()
	repo.AutoMigrate()
	user = domain.NewUserDomain(repo.User)
}
func TestLogin(t *testing.T) {
	search := map[string]interface{}{
		"mobile":   "15829090357",
		"password": "123456",
	}
	fmt.Println(user.Login(search))
}

func TestRegister(t *testing.T) {
	userEntity := new(entity.User)
	userEntity.Password = "123456"
	userEntity.Mobile = uint64(15829090357)
	userEntity.Email = ""
	userEntity.Name = "yino"
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()
	err := user.Add(userEntity)

	fmt.Println("=============")
	fmt.Println(err)
}

func TestCreateAppKey(t *testing.T) {
	fmt.Println(user.CreateAppKey(1, "QA"))
}

func TestAppKeyPage(t *testing.T) {
	fmt.Println(user.AppKeyPage(1, "", 1, 10))
}

func TestAuthAppKey(t *testing.T) {
	fmt.Println(user.AuthAppKey("d3f6cdc561c2216699809d3f17aafa0c", "51fdf32979e3dd778a0b1065216becc5"))
}

func TestDeleteAppKey(t *testing.T) {
	fmt.Println(user.DeleteAppKey(6, 1))
}
