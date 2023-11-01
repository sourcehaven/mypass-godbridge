package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	*gin.Engine
}

func initRouters(engine *gin.Engine) *gin.Engine {
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := engine.Group("/api")
	{
		api.GET("/teapot", routers.IamTeapot)
	}
	return engine
}

func (m *App) Start() {
	host := Cfg.Host
	port := Cfg.Port
	addr := host + ":" + port
	_ = m.Run(addr)
}

func New() *App {
	if Cfg.Env == Production {
		Logger.WithFields(logrus.Fields{
			"topic": "Gin mode",
		}).Info("Entering production environment.")
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	engine = initRouters(engine)
	app := &App{engine}
	return app
}
