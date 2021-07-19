package po

import (
	"time"

	"gorm.io/gorm"
)

// UserAppKeyPo 用户APPKeyPo
type UserAppKeyPo struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	Ak        string         `gorm:"column:app_key;not null;size:255;" json:"app_key"`
	As        string         `gorm:"column:app_secret;not null;size:255;" json:"app_secret"`
	UserID    uint64         `gorm:"type:int(10);not null;" json:"user_id"`
	Type      string         `gorm:"size:20;default:'QA';comment:'类型: QA,....'" json:"type"`
	ReqNum    uint64         `gorm:"type:int(10);not null;default:0;comment:'请求次数'" json:"req_num"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// TableName table name
func (UserAppKeyPo) TableName() string {
	return "user_appkey"
}
