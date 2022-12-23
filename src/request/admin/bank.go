package admin

type AdminBankRequest struct {
	Name string `form:"name" validate:"required" json:"name"`
	Code int    `form:"code" validate:"required,numeric" json:"code"`
}
