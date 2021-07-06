package entity

import (
	"github.com/yino/common"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size;100;not null;" json:"name"`
	Mobile    uint64    `gorm:"size;100;not null;" json:"mobile"`
	Email     string    `gorm:"size;100;not null;" json:"email"`
	Password  string    `gorm:"size;255;not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password = common.MD5(user.Password)
	return nil
}
