package main

import (
	"fmt"
	"module/src/app/api/dicontainer"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Running API...")
	e := echo.New()
	fmt.Println("AAAAAAAAAAAA")
	e.Start(":9000")

	fmt.Println("AAAAAAAAAA")

	e.GET("/artists", dicontainer.GetArtistHandlers().ListArtists)
}
