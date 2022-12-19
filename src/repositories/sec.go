package repositories

import (
	"go-survia/database"
	"go-survia/src/model"
)

type Sec struct{}

func (Sec) All(q string) (b []model.Sec, err error) {
	var secs []model.Sec
	if err = database.DB.Unscoped().Model(&model.Sec{}).Where("name LIKE ?", "%"+q+"%").Order("created_at ASC").Find(&secs).Error; err != nil {
		return secs, err
	}
	return secs, nil
}

func (Sec) FindByID(id string) (r *model.Sec, err error) {
	var sec *model.Sec
	if err = database.DB.Model(&model.Sec{}).First(&sec, "id = ?", id).Error; err != nil {
		return sec, err
	}
	return sec, nil
}

func (Sec) Create(m *model.Sec) (r *model.Sec, err error) {
	if err := database.DB.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (Sec) Patch(m *model.Sec, d interface{}) (r *model.Sec, err error) {
	if err = database.DB.Model(&m).Updates(d).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (Sec) Delete(m *model.Sec) (err error) {
	if err = database.DB.Delete(&m).Error; err != nil {
		return err
	}
	return nil
}