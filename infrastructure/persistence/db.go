package persistence

import (
	"fmt"

	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repositories mysql repo
type Repositories struct {
	User repository.UserRepository
	Qa   repository.QaQuestionRepository
	db   *gorm.DB
}

// NewRepositories new Mysql
func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error() + dsn)
	}
	fmt.Println("MySql连接成功")
	//db.LogMode(true) .
	return &Repositories{
		User: NewUserRepository(db),
		Qa:   NewQaQuestionRepository(db),
		db:   db,
	}, nil
}

// AutoMigrate This migrate all tables
// @return error
func (s *Repositories) AutoMigrate() {
	err := s.db.AutoMigrate(&entity.User{}, &po.UserAppKeyPo{}, &po.QaQuestion{}, &po.UserAppKeyPo{})
	if err != nil {
		panic("migrate fail")
	}
}
