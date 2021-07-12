package aggregate

import (
	"github.com/yino/nlp-controller/domain/repository"
)

// QaFactory: QaQuestion 聚合根
type QaFactory struct {
	UserRepo       repository.UserRepository
	QaQuestionRepo repository.QaQuestionRepository
}

// MatchQuestion 检索问题聚合根
func (factory *QaFactory) MatchQuestion(userId uint64, question string) {

}
