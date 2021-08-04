package po

import (
	"time"

	"github.com/yino/common"

	"gorm.io/gorm"
)

type User struct {
	ID            uint64         `gorm:"primary_key;auto_increment" json:"id"`
	Name          string         `gorm:"size:255;not null;" json:"name"`
	Mobile        uint64         `gorm:"type:bigint(15);not null;index" json:"mobile"`
	Email         string         `gorm:"size:100;not null;" json:"email"`
	Password      string         `gorm:"size:255;not null;" json:"password"`
	Token         string         `gorm:"size:255;not null;" json:"token"`
	QaModelStatus uint8          `gorm:"size:1;not null;default:0;comment:1已训练qa 0未训练;" json:"qa_model_status"`
	TokenExpire   *time.Time     `json:"token_expire"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.Password = common.MD5(user.Password)
	return
}
