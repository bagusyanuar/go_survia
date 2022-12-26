package admin

import "mime/multipart"

type AdminCampaign struct {
	Title            string               `form:"title" validate:"required"`
	Description      string               `form:"description" validate:"required"`
	ShortDescription string               `form:"short_description" validate:"required"`
	StartAt          string               `form:"start_at"`
	FinishAt         string               `form:"finish_at"`
	Background       string               `form:"background" validate:"required"`
	Image            multipart.FileHeader `form:"image" validate:"omitempty"`
}
