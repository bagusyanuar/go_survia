package service

import (
	"go-survia/src/model"
	"go-survia/src/repositories"
	req "go-survia/src/request"
)

type SecQuestion struct {
	repository repositories.SecQuestion
}

func (secQuestion *SecQuestion) Create(r *req.SecQuestion) error {

	lastIndex, e := secQuestion.repository.FindLastIndex()
	if e != nil {
		return e
	}

	entity := model.SecQuestionWithAnswers{
		Question: r.Question,
		IndexOf:  uint(lastIndex),
	}
	if e := secQuestion.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}
