package main

import (
	"net/http"
	"time"

	_wartegHttpHandler "github.com/cpartogi/warteg/module/warteg/handler/http"
	_wartegRepo "github.com/cpartogi/warteg/module/warteg/store"
	_warteg "github.com/cpartogi/warteg/module/warteg/usecase"

	_ "github.com/cpartogi/warteg/docs"
	appInit "github.com/cpartogi/warteg/init"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	log "go.uber.org/zap"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {

	mysqlDb, err := appInit.ConnectToMySqlServer()
	if err != nil {
		log.S().Fatal(err)
	}

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	wartegRepo := _wartegRepo.NewStore(mysqlDb.DB)

	wartegUc := _warteg.NewWartegUsecase(wartegRepo, timeoutContext)

	_wartegHttpHandler.NewWartegHandler(e, wartegUc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
