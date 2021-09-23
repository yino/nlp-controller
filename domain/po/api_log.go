package po

import (
	"time"

	"gorm.io/gorm"
)

const (
	NORMAL  = "NORMAL"
	INVALID = "INVALID"
)

type APILog struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Params    JSON   `gorm:"size:255;not null;type:json;" json:"question"`
	IP        string `gorm:"size:20;not null;" json:"answer"`
	Header    JSON   `gorm:"size:255;not null;type:json;" json:"pid"`
	UserID    uint64 `gorm:"type:int(10);not null;index" json:"user_id"`
	APIType   string `gorm:"size:20;not null;comment:QA;" json:"api_type"`
	APIStatus string `gorm:"size:20;not null;comment:NORMAL,INVALID;" json:"api_status"`
	URL       string `gorm:"size:244;not null;" json:"url"`
	Method    string `gorm:"size:20;not null;" json:"method"`
	CreatedAt string `json:"created_at"`
}

func (APILog) TableName() string {
	return "api_log"
}

func (l *APILog) BeforeCreate(tx *gorm.DB) (err error) {
	l.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return
}

type APILogGroupTime struct {
	Datetime string `json:"datetime"`
	Total    int64  `json:"total"`
}
