package aggregate

import (
	"errors"

	"go.uber.org/zap"

	"github.com/yino/nlp-controller/config/log"

	"github.com/yino/nlp-controller/domain/repository"
	"github.com/yino/nlp-controller/domain/vo"
	"github.com/yino/nlp-controller/infrastructure/api"
)

// QaFactory QaQuestion 聚合根
type QaFactory struct {
	UserRepo       repository.UserRepository
	QaQuestionRepo repository.QaQuestionRepository
}

// MatchQuestion 检索问题聚合根
func (factory *QaFactory) MatchQuestion(uid uint64, question string) (result []vo.QaMatchQuestionItemVo, err error) {
	user, _ := factory.UserRepo.FindAkByUidType(uid, "QA")
	if user.ID == 0 {
		log.Logger.Error("user not fond")
		return result, errors.New("user not fond")
	}
	userInfo, _ := factory.UserRepo.UserInfo(uid)
	if userInfo.QaModelStatus == 0 {
		log.Logger.Error("user not train model")
		return result, errors.New("user not train model")
	}
	qaAPI := api.QaApi{Ak: user.Ak}
	return qaAPI.Match(question)
}

// TrainModel 训练模型
func (factory *QaFactory) TrainModel(uid uint64) error {
	akInfo, _ := factory.UserRepo.FindAkByUidType(uid, "QA")
	if akInfo.ID == 0 {
		log.Logger.Error("user not fond")
		return errors.New("user not fond")
	}
	qaAPI := api.QaApi{Ak: akInfo.Ak}
	err := qaAPI.TrainModel()
	if err != nil {
		log.Logger.Error("api train model fail", zap.Error(err))
		return err
	}

	// update user status
	return factory.UserRepo.UpdateUserQaModel(uid, true)
}

// NewQaFactory 创建QA工厂
func NewQaFactory(userRepo repository.UserRepository, qaRepo repository.QaQuestionRepository) QaFactory {
	return QaFactory{
		UserRepo:       userRepo,
		QaQuestionRepo: qaRepo,
	}
}
