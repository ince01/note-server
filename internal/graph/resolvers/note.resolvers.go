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

func (r *noteResolver) CreatedBy(_ context.Context, obj *model.Note) (*model.User, error) {
	user := &models.User{}

	tx := r.DB.First(user, obj.CreatedBy.ID)

	return &model.User{
		ID:        fmt.Sprint(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     &user.Phone,
		AvatarURL: &user.AvatarUrl,
		CreatedAt: user.CreatedAt,
	}, tx.Error
}

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

type noteResolver struct{ *Resolver }
