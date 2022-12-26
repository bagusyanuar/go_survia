package service

import (
	"go-survia/src/lib"
	"go-survia/src/repositories"
	adminRequest "go-survia/src/request/admin"
	"mime/multipart"
)

type Campaign struct {
	repository repositories.Campaign
}

func (campaign *Campaign) Create(request *adminRequest.AdminCampaign) (d interface{}, f *multipart.FileHeader, fname *string, err error) {
	// var startAt *datatypes.Date
	// var finishAt *datatypes.Date
	var image *string

	messages, e := lib.ValidateRequest(request)
	if e != nil {
		return messages, nil, image, lib.ErrBadRequest
	}

	// if request.StartAt != "" {
	// 	tmp, e := time.Parse("2006-01-02", request.StartAt)
	// 	if e != nil {
	// 		return nil, nil, image, e
	// 	}
	// 	startAt = (*datatypes.Date)(&tmp)
	// }

	// if request.FinishAt != "" {
	// 	tmp, e := time.Parse("2006-01-02", request.FinishAt)
	// 	if e != nil {
	// 		return nil, nil, image, e
	// 	}
	// 	finishAt = (*datatypes.Date)(&tmp)
	// }

	file := request.Image

	// if file != nil {
	// 	ext := filepath.Ext(file.Filename)
	// 	fileName := "assets/campaigns/" + uuid.New().String() + ext
	// 	image = &fileName
	// }

	// entity := model.Campaign{
	// 	Title:            request.Title,
	// 	Description:      request.Description,
	// 	ShortDescription: request.ShortDescription,
	// 	StartAt:          startAt,
	// 	FinishAt:         finishAt,
	// 	Background:       request.Background,
	// 	Image:            image,
	// }
	// e = campaign.repository.Create(&entity)
	// if e != nil {
	// 	return nil, file, image, e
	// }
	return nil, &file, image, nil
}
