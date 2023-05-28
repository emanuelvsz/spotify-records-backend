package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func loadApiRoutes(group *echo.Group) {
	group.GET("/*", echoSwagger.WrapHandler)
}

func (instance *router) Load(group *echo.Group) {
	loadApiRoutes(group)
}
