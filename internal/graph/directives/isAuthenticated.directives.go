// package directives

// import (
// 	"context"
// 	"fmt"

// 	"github.com/99designs/gqlgen/graphql"
// )

// func (d *Directives) isAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
// 	// if !getCurrentUser(ctx).HasRole(role) {
// 		// block calling the next resolver
// 		return nil, fmt.Errorf("Access denied")
// 	}

// 	// or let it pass through
// 	return next(ctx)
// }
