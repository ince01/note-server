package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ince01/note-server/internal/graph/generated"
	"github.com/ince01/note-server/internal/graph/model"
	"github.com/ince01/note-server/internal/orm/models"
)

func (r *noteResolver) Children(ctx context.Context, obj *model.Note) (*model.Note, error) {
	return obj, nil
}

func (r *noteResolver) CreatedBy(ctx context.Context, obj *model.Note) (*model.User, error) {
	var user models.User

	tx := r.DB.First(&user, obj.CreatedBy)

	if tx.Error != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &model.User{
		ID:        fmt.Sprint(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarUrl,
		CreatedAt: user.CreatedAt,
	}, nil
}

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

type noteResolver struct{ *Resolver }
