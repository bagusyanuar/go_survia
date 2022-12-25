package admin

type AdminBank struct {
	Name string `form:"name" validate:"required" json:"name"`
	Code int    `form:"code" validate:"required,numeric" json:"code"`
}
