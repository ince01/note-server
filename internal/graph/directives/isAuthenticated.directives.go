package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ince01/note-server/internal/auth"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := auth.ForContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("access denied")
	}

	return next(ctx)
}
