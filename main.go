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

	restricted := e.Group("/restricted")
	restricted.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		ContextKey: string(authContextKey),
		SigningKey: []byte("secret"),
	}))
	restricted.POST("/graphql", func(c echo.Context) error {
		h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}),
			handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
				ctx, err := getCtxWithJWTCtxFromEchoCtx(ctx, c)
				if err != nil {
					return nil, err
				}
				return next(ctx)
			}),
		)
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}

func getCtxWithJWTCtxFromEchoCtx(ctx context.Context, c echo.Context) (context.Context, error) {
	token, ok := c.Get(string(authContextKey)).(*jwt.Token)
	if ok == false {
		return nil, errors.New("auth_context_not_found")
	}
	return context.WithValue(ctx, authContextKey, token), nil
}
