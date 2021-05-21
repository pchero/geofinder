package apiv1

import (
	"github.com/gin-gonic/gin"

	"github.com/pchero/geofinder/api/v1.0/finds"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")

	// v1.0
	finds.ApplyRoutes(v1)
}
