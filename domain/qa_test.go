package domain_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/yino/nlp-controller/domain/entity"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/infrastructure/persistence"
)

var qa domain.Qa

func init() {
	os.Setenv("env", "test")
	config.GetConf()
	repo, _ := persistence.NewRepositories(config.Conf.MySql.User, config.Conf.MySql.Password, config.Conf.MySql.Port, config.Conf.MySql.Host, config.Conf.MySql.Db)
	log.InitLogger()
	//repo.AutoMigrate()
	qa = domain.NewQaDomain(repo.Qa)
}
func TestQa_Page(t *testing.T) {
	search := make(map[string]interface{})
	list, total, err := qa.Page(2, 10, search)
	if err != nil {
		fmt.Println("获取page 失败", err)
	}
	fmt.Println("total", total)
	fmt.Println("len", len(list))
	fmt.Println(list)
}
func TestQa_Add(t *testing.T) {

	wg := sync.WaitGroup{}
	for i := 0; i <= 500; i++ {
		wg.Add(1)
		entityQa := new(entity.QaQuestion)
		entityQa.Pid = 0
		entityQa.Question = fmt.Sprintf("测试%d", i)
		entityQa.Answer = fmt.Sprintf("测试数据%d", i)
		entityQa.Type = 1
		entityQa.UserId = 1

		go func() {
			err := qa.Add(entityQa)
			fmt.Println(err)
			wg.Done()
		}()
	}
	wg.Wait()

}
