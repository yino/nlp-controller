package application_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yino/nlp-controller/application"
	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/infrastructure/persistence"
)
var userApp application.UserApp
func init(){
	config.GetConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	userApp = application.NewUserApp(repo.User)
}
func TestLogin(t *testing.T) {
	search := map[string]interface{}{
		"mobile":   "15829090357",
		"password": "123456",
	}
	fmt.Println(userApp.Login(search))
}

func TestRegister(t *testing.T) {
	userEntity := new(entity.User)
	userEntity.Password = "123456"
	userEntity.Mobile = uint64(15829090357)
	userEntity.Email = ""
	userEntity.Name = "yino"
	userEntity.CreatedAt = time.Now()
	userEntity.UpdatedAt = time.Now()
	fmt.Println(userApp.Add(userEntity))
}
