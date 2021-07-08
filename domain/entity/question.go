package entity

type Question struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
}
