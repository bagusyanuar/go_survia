package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	request "go-survia/src/request/admin"
	adminResponse "go-survia/src/response/admin"

	"github.com/google/uuid"
)

type Category struct{}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var apiCategories []CategoryResponse

//admin
func (Category) All(q string) (b []adminResponse.APICategoryResponse, err error) {
	var categories []adminResponse.APICategoryResponse
	if err = database.DB.Unscoped().Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (Category) FindByID(id string) (r *adminResponse.APICategoryResponse, err error) {
	var category *adminResponse.APICategoryResponse
	if err = database.DB.Model(&model.Category{}).First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (Category) Create(request *request.AdminCategoryRequest) (r *model.Category, err error) {
	m := model.Category{
		Name: request.Name,
	}
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (Category) Patch(id string, d interface{}) (err error) {

	var category *model.Category
	if err = database.DB.Debug().Model(&category).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Category) Delete(id string) (err error) {
	var category *model.Category
	// if err = database.DB.Model(&model.Category{}).First(&category, "id = ?", id).Error; err != nil {
	// 	return err
	// }
	if err = database.DB.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (Category) GetCategories(q string) (r []CategoryResponse, err error) {
	if err = database.DB.Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&apiCategories).Error; err != nil {
		return apiCategories, err
	}
	return apiCategories, nil
}
