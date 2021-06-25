package repository

type QuestionRepository interface {
	GetPage(page, limit int64,search map[string]string)
}
