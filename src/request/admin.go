package request

type Category struct {
	Name string `form:"name" validate:"required" json:"name"`
}

type Bank struct {
	Name string `form:"name" validate:"required" json:"name"`
	Code int    `form:"code" validate:"required,numeric" json:"code"`
}

type Province struct {
	Code int    `form:"code" validate:"required" json:"code"`
	Name string `form:"name" validate:"required" json:"name"`
}

type City struct {
	ProvinceID string `form:"province_id" validate:"required" json:"province_id"`
	Code       int    `form:"code" validate:"required" json:"code"`
	Name       string `form:"name" validate:"required" json:"name"`
}

type Job struct {
	Name string `form:"name" validate:"required" json:"name"`
}

type Sec struct {
	Name   string `form:"name" validate:"required"`
	Bottom int    `form:"bottom" validate:"required"`
	Top    int    `form:"top" validate:"required"`
}

type SecQuestion struct {
	Question string `json:"question" validate:"required"`
	Answers  []struct {
		Answer string `json:"answer" validate:"required"`
		Score  uint   `json:"score" validate:"required"`
	} `json:"answers" validate:"required"`
}
