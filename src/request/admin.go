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
