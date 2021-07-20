package domain_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/yino/nlp-controller/config"
	"github.com/yino/nlp-controller/config/log"
	"github.com/yino/nlp-controller/domain"
	"github.com/yino/nlp-controller/domain/entity"
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
	vo, err := qa.GetMasterQuestionPage(2, 10, search)
	if err != nil {
		fmt.Println("获取page 失败", err)
	}
	fmt.Println(vo)
}

func TestQa_BashAdd(t *testing.T) {

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
			err := qa.AddMaster(entityQa)
			fmt.Println(err)
			wg.Done()
		}()
	}
	wg.Wait()

}
func TestQa_Add(t *testing.T) {

	entityQa := new(entity.QaQuestion)
	entityQa.Pid = 0
	entityQa.Question = "测试10086"
	entityQa.Answer = "测试10086"
	entityQa.Type = 1
	entityQa.UserId = 1
	fmt.Println(entityQa)
	var slaveQuestion []entity.QaQuestion

	entitySlaveQa := entity.QaQuestion{}
	entitySlaveQa.Pid = 0
	entitySlaveQa.Question = "测试1008611"
	entitySlaveQa.Answer = "测试1008611"
	entitySlaveQa.Type = 1
	entitySlaveQa.UserId = 1
	slaveQuestion = append(slaveQuestion, entitySlaveQa)
	slaveQuestion = append(slaveQuestion, entitySlaveQa)
	slaveQuestion = append(slaveQuestion, entitySlaveQa)
	err := qa.Add(entityQa, slaveQuestion)
	fmt.Println(" error ", err)
}
func TestQa_MasterEdit(t *testing.T) {
	entityQa := new(entity.QaQuestion)
	entityQa.Pid = 0
	entityQa.Question = "测试10001"
	entityQa.Answer = "测试10001"
	entityQa.Type = 1
	entityQa.UserId = 1
	entityQa.ID = 380
	fmt.Println(qa.EditMaster(entityQa))
}

func TestQa_Edit(t *testing.T) {
	entityQa := new(entity.QaQuestion)
	entityQa.Pid = 0
	entityQa.Question = "test11"
	entityQa.Answer = "test11"
	entityQa.Type = 1
	entityQa.UserId = 1
	entityQa.ID = 8052

	var slaveQuestion []entity.QaQuestion
	var testQaEnt entity.QaQuestion
	testQaEnt.ID = 10868
	testQaEnt.Question = "test slave 32"
	testQaEnt.Answer = "test slave 32"
	testQaEnt.Pid = 8052
	slaveQuestion = append(slaveQuestion, testQaEnt)

	testQaEnt.ID = 10869
	testQaEnt.Question = "test slave 22"
	testQaEnt.Answer = "test slave 22"
	testQaEnt.Pid = 8052
	slaveQuestion = append(slaveQuestion, testQaEnt)

	testQaEnt.ID = 10870
	testQaEnt.Question = "test slave 12"
	testQaEnt.Answer = "test slave 12"
	testQaEnt.Pid = 8052
	slaveQuestion = append(slaveQuestion, testQaEnt)

	testQaEnt.ID = 10871
	testQaEnt.Question = "test slave 02"
	testQaEnt.Answer = "test slave 02"
	testQaEnt.Pid = 8052
	slaveQuestion = append(slaveQuestion, testQaEnt)

	err := qa.Edit(uint64(1), entityQa, slaveQuestion)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("修改成功")
	}
}
func TestQa_FindInfo(t *testing.T) {
	fmt.Println(qa.FindInfo(2566))
}

func TestQa_Delete(t *testing.T) {
	fmt.Println(qa.Delete(uint64(1), 379))
}
