package persistence

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm/logger"

	"github.com/yino/nlp-controller/domain/po"
	"github.com/yino/nlp-controller/domain/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repositories mysql repo
type Repositories struct {
	User   repository.UserRepository
	Qa     repository.QaQuestionRepository
	APILog repository.APILogRepository
	db     *gorm.DB
}

var (
	dbFactory Repositories
	once      sync.Once
)

// NewRepositories new Mysql
func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {

	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,   // 慢 SQL 阈值
				LogLevel:      logger.Silent, // Log level
				Colorful:      false,         // 禁用彩色打印
			},
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic("数据库连接失败:" + err.Error() + dsn)
		}
		fmt.Println("MySql连接成功")
		//db.LogMode(true) .
		dbFactory = Repositories{
			User:   NewUserRepository(db),
			Qa:     NewQaQuestionRepository(db),
			APILog: NewLogRepository(db),
			db:     db,
		}
	})

	return &dbFactory, nil
}

// AutoMigrate This migrate all tables
// @return error
func (s *Repositories) AutoMigrate() {
	err := s.db.AutoMigrate(
		//&po.User{},
		//&po.UserAppKeyPo{},
		//&po.QaQuestion{},
		&po.APILog{},
	)
	if err != nil {
		panic("migrate fail")
	}
}
