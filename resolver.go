package main

import "context"

// Resolver defines struct
type Resolver struct{}

// Query returns queryResolver
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (User, error) {
	panic("not implemented")
}
