package aggregate

import (
	"github.com/yino/nlp-controller/domain/entity"
	"github.com/yino/nlp-controller/domain/repository"
)

// QaQuestionFactory: QaQuestion 聚合根
type QaQuestionFactory struct {
	UserRepo       repository.UserRepository
	QaQuestionRepo repository.QaQuestionRepository
}

// NewQuestion: 创建question
func (factory *QaQuestionFactory) NewQuestion(userId uint64, question []entity.QaQuestion) error {
	return nil
}
