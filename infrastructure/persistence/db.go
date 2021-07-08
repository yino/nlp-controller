package persistence

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" //这个一定要引入哦！
	"gorm.io/gorm"
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error() + dsn)
		return nil, err
	}
	//db.LogMode(true) .
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

//This migrate all tables
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.User{})
}
