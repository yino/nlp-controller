package po

import (
	"time"
)

const (
	NORMAL  = "NORMAL"
	INVALID = "INVALID"
)

type APILog struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Params    JSON      `gorm:"size:255;not null;type:json;" json:"question"`
	IP        string    `gorm:"size:20;not null;" json:"answer"`
	Header    JSON      `gorm:"type:int(10);default:0;index" json:"pid"`
	UserID    uint64    `gorm:"type:int(10);not null;index" json:"user_id"`
	APIType   string    `gorm:"size:20;not null;comment:QA;" json:"api_type"`
	APIStatus string    `gorm:"size:20;not null;comment:NORMAL,INVALID;" json:"api_status"`
	URL       string    `gorm:"size:244;not null;" json:"url"`
	Method    string    `gorm:"size:20;not null;" json:"method"`
	CreatedAt time.Time `json:"created_at"`
}

func (APILog) TableName() string {
	return "api_log"
}
