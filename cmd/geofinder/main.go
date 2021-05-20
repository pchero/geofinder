package main

import (
	"flag"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/pchero/geofinder/api"
	"github.com/pchero/geofinder/api/models/common"
	"github.com/pchero/geofinder/pkg/servicehandler"
)

// @title VoIPBIN project API
// @version 1.0
// @description RESTful API documents for VoIPBIN project.
// @termsOfService http://swagger.io/terms/

// @contact.name VoIPBIN Project
// @contact.email pchero21@gmail.com

// @host api.voipbin.net
// @BasePath
func main() {

	// create servicehandler
	serviceHandler := servicehandler.NewServiceHandler()

	app := gin.Default()

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// docs
	app.Static("/docs", "docsdev/build/html")

	// CORS setting
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// inject servicehandler
	app.Use(func(c *gin.Context) {
		c.Set(common.OBJServiceHandler, serviceHandler)
		c.Next()
	})

	// apply api router
	api.ApplyRoutes(app)

	logrus.Debug("Starting the api service.")
	app.Run(":8080")
}

func init() {
	flag.Parse()

	// init log
	logrus.SetFormatter(joonix.NewFormatter())
	logrus.SetLevel(logrus.DebugLevel)
}
