package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":3000"))
}
