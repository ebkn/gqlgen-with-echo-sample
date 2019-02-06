package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "username" && password == "password" {
		claims := &jwt.StandardClaims{
			Subject:   username,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedStr, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		return c.String(http.StatusOK, signedStr)
	}
	return c.NoContent(http.StatusUnauthorized)
}
