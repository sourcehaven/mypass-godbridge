package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mypass-godbridge/pkg/routers"
	"strings"
)

// TODO: this should be called before every other stuff. Could this be forced somehow?
func initEnv() {
	cfg := Config{}
	env := cfg.Env()
	env = strings.ToLower(env)
	if env != "" {
		_ = godotenv.Load(".env." + env)
	}
	_ = godotenv.Load()
}

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
	cfg := Config{}
	host := cfg.Host()
	port := cfg.Port()
	addr := host + ":" + port
	_ = m.Run(addr)
}

func CreateApp() App {
	initEnv()
	engine := initEngine()
	app := App{engine}
	return app
}
