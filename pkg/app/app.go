package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mypass-godbridge/pkg/routers"
)

type App struct {
	*gin.Engine
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/teapot", routers.IamTeapot)
	}
	return router
}

func initEngine() *gin.Engine {
	router := initRouter()
	return router
}

func (m *App) Start() {
	host := Cfg.Host
	port := Cfg.Port
	addr := host + ":" + port
	_ = m.Run(addr)
}

func CreateApp() *App {
	engine := initEngine()
	app := &App{engine}
	return app
}
