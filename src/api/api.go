package api

import (
	"github.com/labstack/echo/v4"
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
