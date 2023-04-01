package api

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API interface {
	Serve()
	loadRoutes()
}

type Options struct{}

type api struct {
	options      *Options
	group        *echo.Group
	echoInstance *echo.Echo
}

// NewAPI
// @title COLABORA RTP API
// @version 1.0
// @description Gestão de projetos do DIT
// @contact.name DIT - IFAL
// @contact.email evs10@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey
// @in header
// @name Authorization
func NewAPI(options *Options) API {
	echoInstance := echo.New()
	return &api{options, echoInstance.Group("/api"), echoInstance}
}

func (instance *api) Serve() {
	instance.echoInstance.Use(middleware.Logger())
	instance.echoInstance.Use(middleware.Recover())
	instance.echoInstance.Use(instance.getCORSSettings())
	instance.loadRoutes()
	address := getServerAddress()
	instance.echoInstance.Logger.Fatal(instance.echoInstance.Start(address))
}

func (instance *api) loadRoutes() {
	router := router.New()
	router.Load(instance.group)
}

func (instance *api) getCORSSettings() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:         middlewares.OriginInspectSkipper,
		AllowOriginFunc: middlewares.VerifyOrigin,
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})
}

func getServerAddress() string {
	host := utils.GetenvWithDefault("SERVER_HOST", "0.0.0.0")
	port := utils.GetenvWithDefault("SERVER_PORT", "8000")
	return fmt.Sprintf("%s:%s", host, port)
}

func Logger() zerolog.Logger {
	return log.With().Str("layer", "api").Logger()
}