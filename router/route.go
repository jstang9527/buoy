package router

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/jstang9527/buoy/controller"
	"github.com/jstang9527/buoy/docs"
	"github.com/jstang9527/buoy/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter ...
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = lib.GetStringConf("base.swagger.version")
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	// docs.SwaggerInfo.Schemes = strings.Split(lib.GetStringConf("base.swagger.schemes"), ",")
	docs.SwaggerInfo.Schemes = lib.GetStringSliceConf("base.swagger.schemes")

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	demoRouter := router.Group("/demo")
	demoRouter.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		// middleware.IPAuthMiddleware(),
	)
	controller.DemoRegister(demoRouter)
	return router
}
