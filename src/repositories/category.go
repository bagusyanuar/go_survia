package repositories

import (
	"go-survia/database"
	"go-survia/src/model"

	"github.com/gofrs/uuid"
)

type Category struct{}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var categories []model.Category
var category *model.Category

var apiCategories []CategoryResponse

//admin
func (Category) All(q string) (b []model.Category, err error) {
	if err = database.DB.Unscoped().Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (Category) FindByID(id string) (r *model.Category, err error) {
	if err = database.DB.Model(&model.Category{}).First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (Category) Create(m *model.Category) (r *model.Category, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Category) Patch(m *model.Category, d interface{}) (r *model.Category, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Category) Delete(m *model.Category) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
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
