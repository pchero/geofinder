package finds

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/pchero/geofinder/api/models/common"
	"github.com/pchero/geofinder/api/models/request"
	"github.com/pchero/geofinder/api/models/response"
	"github.com/pchero/geofinder/pkg/servicehandler"
)

// findsPOST handles POST /finds request.
// It checks a given address's country has listed in the given list.
// @Summary It checks a given address's country has listed in the given list.
// @Description It checks a given address's country has listed in the given list.
// @Produce json
// @Param data body request.BodyFindsPOST true "Data"
// @Success 200 {object} response.BodyFindsPOST
// @Router /v1.0/finds [post]
func findsPOST(c *gin.Context) {

	var body request.BodyFindsPOST
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check the geo
	serviceHandler := c.MustGet(common.OBJServiceHandler).(servicehandler.ServiceHandler)
	tmp, err := serviceHandler.GeoFind(body.Address, body.Countries)
	if err != nil {
		logrus.Errorf("Could not create a flow. err: %v", err)
		c.AbortWithStatus(400)
		return
	}

	res := response.BodyFindsPOST{
		IsListed: tmp,
	}

	c.JSON(200, res)
	return
}
