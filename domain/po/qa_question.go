package po

import (
	"time"

	"gorm.io/gorm"
)

type QaQuestion struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	Question  string         `gorm:"size:255;not null;" json:"question"`
	Answer    string         `gorm:"size:255;not null;" json:"answer"`
	Pid       uint64         `gorm:"type:int(10);default:0;index" json:"pid"`
	UserId    uint64         `gorm:"type:int(10);not null;index" json:"user_id"`
	Type      uint64         `gorm:"default:1;size:4;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (QaQuestion) TableName() string {
	return "qa_questions"
}
