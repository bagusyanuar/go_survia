package admin

type AdminProvince struct {
	Code int    `form:"code" validate:"required" json:"code"`
	Name string `form:"name" validate:"required" json:"name"`
}
