package routes

import (
	"module/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadAuthRoutes(group *echo.Group) {
	authGroup := group.Group("/auth")

	authHandlers := dicontainer.GetAuthHandlers()

	authGroup.POST("/login", authHandlers.Login)
}
