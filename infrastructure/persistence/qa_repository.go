package persistence

import (
	"time"

	"github.com/yino/nlp-controller/domain/po"

	"gorm.io/gorm"
)

// QaQuestionRepo qa repo infra
type QaQuestionRepo struct {
	db *gorm.DB
}

// NewQaQuestionRepository qa test
func NewQaQuestionRepository(db *gorm.DB) *QaQuestionRepo {
	return &QaQuestionRepo{db: db}
}

// Page page
func (obj *QaQuestionRepo) Page(page, limit int64, search map[string]interface{}) (list []po.QaQuestion, total int64, err error) {
	countDb := obj.db
	db := obj.db
	if userID, ok := search["user_id"]; ok {
		countDb.Where("user_id = ?", userID)
		db.Where("user_id = ?", userID)
	}

	if page > 0 {
		page = page - 1
	}
	countDb.Model(&po.QaQuestion{}).Count(&total)
	if total == 0 {
		return
	}
	db.Limit(int(limit)).Offset(int(page * limit)).Order("id desc").Find(&list)
	return
}

// Add add
func (obj *QaQuestionRepo) Add(question *po.QaQuestion) error {
	return obj.db.Save(question).Error

}

// Edit edit
func (obj *QaQuestionRepo) Edit(question *po.QaQuestion) error {
	question.UpdatedAt = time.Now()
	return obj.db.Where("id = ?", question.ID).Save(question).Error
}

//Delete delete
func (obj *QaQuestionRepo) Delete(id uint64) error {

	return obj.db.Where("id = ?", id).Delete(&po.QaQuestion{}).Error
}

// BatchInsert 批量插入
func (obj *QaQuestionRepo) BatchInsert([]po.QaQuestion) error {
	return nil
}

// FindInfo 查询根据ID详情
func (obj *QaQuestionRepo) FindInfo(id uint64) (*po.QaQuestion, error) {
	poQuestion := new(po.QaQuestion)
	err := obj.db.Where(po.QaQuestion{ID: id}).Find(poQuestion).Error
	return poQuestion, err
}
