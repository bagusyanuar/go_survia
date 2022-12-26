package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
	adminResponse "go-survia/src/response/admin"
)

type Category struct{}

//admin
func (Category) All(q string) (b []adminResponse.APICategory, err error) {
	var categories []adminResponse.APICategory
	if err = database.DB.Unscoped().Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (Category) FindByID(id string) (r *adminResponse.APICategory, err error) {
	var category *adminResponse.APICategory
	if err = database.DB.Model(&model.Category{}).First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (Category) Create(entity *model.Category) error {
	if err := database.DB.Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (Category) Patch(id string, d interface{}) error {
	if err := database.DB.Debug().Model(&model.Category{}).Where("id = ?", id).Updates(d).Error; err != nil {
		return err
	}
	return nil
}

func (Category) Delete(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}

// func (Category) GetCategories(q string) (r []CategoryResponse, err error) {
// 	if err = database.DB.Model(&model.Category{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&apiCategories).Error; err != nil {
// 		return apiCategories, err
// 	}
// 	return apiCategories, nil
// }
