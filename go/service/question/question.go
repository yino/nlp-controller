package question

type QuestionCount interface {
	Total() int64
}

type nameA struct{}

func (n *nameA) Total() int64 {}

type nameB struct{}

func (n *nameB) Total() int64 {}
