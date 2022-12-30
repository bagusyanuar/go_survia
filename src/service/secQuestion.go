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

	var answers []model.SecAnswer

	for k, v := range r.Answers {
		tmpAnswers := model.SecAnswer{
			Answer:  v.Answer,
			Score:   int(v.Score),
			IndexOf: uint((k + 1)),
		}
		answers = append(answers, tmpAnswers)
	}
	entity := model.SecQuestionWithAnswers{
		Question: r.Question,
		IndexOf:  uint(lastIndex),
		Answers: answers,
	}
	if e := secQuestion.repository.Create(&entity); e != nil {
		return e
	}
	return nil
}

func (secQuestion *SecQuestion) FindAll(q string) (d []model.SecQuestionWithAnswers, err error) {
	return secQuestion.repository.All(q)
}

func (secQuestion *SecQuestion) FindByID(id string) (d *model.SecQuestionWithAnswers, err error) {
	data, e := secQuestion.repository.FindByID(id)
	if e != nil {
		return nil, e
	}
	return data, nil
}