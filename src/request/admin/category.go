package admin

type AdminCategory struct {
	Name string `form:"name" validate:"required" json:"name"`
}
