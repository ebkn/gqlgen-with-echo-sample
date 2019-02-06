package main

import (
	"context"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// Resolver defines struct
type Resolver struct{}

// Query returns queryResolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (User, error) {
	username, err := getJWTSubjectFromCtx(ctx)
	if err != nil {
		return User{}, err
	}
	return User{Username: username}, nil
}

func getJWTSubjectFromCtx(ctx context.Context) (string, error) {
	token, ok := ctx.Value(authContextKey).(*jwt.Token)
	if ok == false {
		return "", errors.New("auth_context_not_found")
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Subject, nil
}
