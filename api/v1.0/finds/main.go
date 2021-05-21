package finds

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	finds := r.Group("/finds")

	finds.POST("", findsPOST)
}
