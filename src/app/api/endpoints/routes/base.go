package routes

import (
	_ "module/src/app/api/docs"

	"github.com/labstack/echo/v4"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func (instance *router) Load(group *echo.Group) {
	loadApiRoutes(group)
	loadArtistRoutes(group)
}
