package persistence

import (
	"fmt"
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
		countDb = countDb.Where("user_id = ?", userID)
		db = db.Where("user_id = ?", userID)
	}
	if pid, ok := search["pid"]; ok {
		countDb = countDb.Where("pid = ?", pid)
		db = db.Where("pid = ?", pid)
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

// AddMaster add master question
func (obj *QaQuestionRepo) AddMaster(question *po.QaQuestion) (uint64, error) {
	err := obj.db.Save(question).Error
	return question.ID, err
}

// EditMaster edit maser question
func (obj *QaQuestionRepo) EditMaster(question *po.QaQuestion) error {
	question.UpdatedAt = time.Now()
	return obj.db.Where("id = ?", question.ID).Save(question).Error
}

//Delete delete
func (obj *QaQuestionRepo) Delete(id uint64) error {
	return obj.db.Transaction(func(tx *gorm.DB) error {
		deletQa := po.QaQuestion{}
		if err := tx.Where("id = ?", id).Delete(&deletQa).Error; err != nil {
			return err
		}
		if err := tx.Where("pid = ?", id).Delete(&deletQa).Error; err != nil {
			return err
		}
		return nil
	})
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

// GetSlaveList 获取相似问题
func (obj *QaQuestionRepo) GetSlaveList(pid uint64) ([]po.QaQuestion, error) {
	var list []po.QaQuestion
	err := obj.db.Where("pid = ?", pid).Find(&list).Error
	return list, err
}

// Add add  question
func (obj *QaQuestionRepo) Add(question *po.QaQuestion, slaveQuestion []po.QaQuestion) error {
	return obj.db.Transaction(func(tx *gorm.DB) error {
		// master add
		if err := obj.db.Save(question).Error; err != nil {
			return err
		}

		// slave add
		var insertSlaveData []po.QaQuestion
		for _, slaveItem := range slaveQuestion {
			item := slaveItem
			item.Pid = question.ID
			insertSlaveData = append(insertSlaveData, item)
		}
		if len(insertSlaveData) > 0 {
			if err := obj.db.Create(&insertSlaveData).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// Edit edit  question
func (obj *QaQuestionRepo) Edit(question *po.QaQuestion, slaveQuestion []po.QaQuestion) error {
	//更新、添加、删除的slave question
	var updateQuestionList, insertQuestionList, questionList []po.QaQuestion
	var deleteQuestionList []uint64
	obj.db.Where("pid = ?", question.ID).Where("user_id = ?", question.UserId).Find(&questionList)
	if len(slaveQuestion) > 0 {
		fmt.Println("len(questionList)", len(questionList))
		fmt.Println("slaveQuestion", slaveQuestion)
		// 如果没有slave question list的时候 则全部为插入操作
		if len(questionList) == 0 {
			insertQuestionList = slaveQuestion
		} else {
			slaveQuestionMap := make(map[uint64]po.QaQuestion)
			// 将questionList 转换为 map[id]struct 牺牲空间，减少执行次数
			// 拼接 add  slice
			for _, questionPo := range slaveQuestion {
				if questionPo.ID == 0 {
					// insert
					insertQuestionList = append(insertQuestionList, questionPo)
					continue
				}
				slaveQuestionMap[questionPo.ID] = questionPo
			}
			// 拼接 delete and update question
			for _, questionPo := range questionList {
				// delete
				submitQuestion, ok := slaveQuestionMap[questionPo.ID]
				if !ok {
					deleteQuestionList = append(deleteQuestionList, questionPo.ID)
					continue
				}
				// update
				if questionPo.Question != submitQuestion.Question {
					updateQuestionList = append(updateQuestionList, submitQuestion)
				}
			}
		}
	} else {
		for _, slaveQuestionItem := range questionList {
			deleteQuestionList = append(deleteQuestionList, slaveQuestionItem.ID)
		}
	}

	fmt.Println("insert", insertQuestionList)
	fmt.Println("update", updateQuestionList)
	fmt.Println("delete", deleteQuestionList)
	return obj.db.Transaction(func(tx *gorm.DB) error {

		question.UpdatedAt = time.Now()
		if err := tx.Where("id = ?", question.ID).Save(question).Error; err != nil {
			return err
		}

		// insert
		if len(insertQuestionList) > 0 {
			if err := tx.Create(&insertQuestionList).Error; err != nil {
				return err
			}
		}

		// delete
		if len(deleteQuestionList) > 0 {
			deleteQaPo := po.QaQuestion{}
			if err := tx.Where("id in ?", deleteQuestionList).Where("user_id = ?", question.UserId).Delete(&deleteQaPo).Error; err != nil {
				return err
			}
		}

		// update
		if len(updateQuestionList) > 0 {
			for _, updateQuestionItem := range updateQuestionList {
				if err := tx.Model(&updateQuestionItem).Where("id = ?", updateQuestionItem.ID).Where("user_id = ?", question.UserId).Updates(po.QaQuestion{Question: updateQuestionItem.Question, Answer: question.Answer}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}
