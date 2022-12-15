package admin

import (
	"go-survia/database"
	"go-survia/src/model"

	"github.com/gin-gonic/gin"
)


type Campaign struct{}

var listCampaign []model.Campaign
var campaign *model.Campaign

func (Campaign) Index(c *gin.Context)  {
	
}

//logical
func findAllCampaign(q string) (r []model.Campaign, err error) {
	//unscoped for show deleted item
	if err = database.DB.Unscoped().Model(&model.Campaign{}).Where("name LIKE ?", "%"+q+"%").Find(&listCampaign).Error; err != nil {
		return listCampaign, err
	}
	return listCampaign, nil
}