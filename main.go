package main

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// AuthContextKeyType type of authContextKey
type AuthContextKeyType string

const authContextKey AuthContextKeyType = "auth"

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.POST("/login", login)

	e.GET("/playground", echo.WrapHandler(handler.Playground("GraphQL playground", "/api/graphql")))

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		ContextKey: string(authContextKey),
		SigningKey: []byte("secret"),
	}))
	api.POST("/graphql", func(c echo.Context) error {
		h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}),
			handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
				token, ok := c.Get(string(authContextKey)).(*jwt.Token)
				if ok == false {
					return nil, errors.New("auth_context_not_found")
				}
				ctx = context.WithValue(ctx, authContextKey, token)
				return next(ctx)
			}),
		)
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}
