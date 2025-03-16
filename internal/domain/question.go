package domain

import (
	"helloworld/internal/data/question/model"
	"time"
)

type QuestionBase struct {
	ID         int64
	Type       string
	Content    string
	ImagePath  string
	Difficulty float64
	Knowledge  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type QuestionChoice struct {
	QuestionBase
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Answer  string
}

func (u *QuestionChoice) padding(base *model.QuestionBase,
	choice *model.QuestionChoice) {
	u.ID = base.ID
	u.Type = base.Type
	u.Content = base.Content
	u.ImagePath = base.ImagePath
	u.Difficulty = float64(base.Difficulty)
	//u.Knowledge = base.KnowledgeID
	u.OptionA = choice.OptionA
	u.OptionB = choice.OptionB
	u.OptionC = choice.OptionC
	u.OptionD = choice.OptionD
	u.Answer = choice.Answer
}
