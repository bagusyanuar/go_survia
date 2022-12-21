package admin

type AdminCategoryRequest struct {
	Name string `form:"name" validate:"required" json:"name"`
}