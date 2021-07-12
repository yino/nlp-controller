package po

import (
	"time"

	"gorm.io/gorm"
)

// 用户APPKeyPo
type UserAppKeyPo struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	AppKey    string         `gorm:"not null;size:255;" json:"app_key"`
	AppSecret string         `gorm:"not null;size:255;" json:"app_secret"`
	UserId    uint64         `gorm:"type:int(10);not null;foreignKey:UserId;index" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (u *UserAppKeyPo) TableName() string {
	return "user_appkey"
}
