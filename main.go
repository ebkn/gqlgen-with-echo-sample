package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const authContextKey = "auth"

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.POST("/login", login)

	restricted := e.Group("/restricted")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		ContextKey: authContextKey,
		SigningKey: []byte("secret"),
	}))
	e.Logger.Fatal(e.Start(":3000"))
}
