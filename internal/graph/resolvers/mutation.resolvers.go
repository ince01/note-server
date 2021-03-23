package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ince01/note-server/internal/graph/generated"
	"github.com/ince01/note-server/internal/graph/model"
	"github.com/ince01/note-server/internal/orm/models"
)

func (r *mutationResolver) NoteCreate(ctx context.Context, note model.NoteInput) (*model.Note, error) {
	tx := r.DB.First(&model.User{ID: note.CreatedBy})

	if tx.Error != nil {
		return nil, tx.Error
	}

	u64, _ := strconv.ParseUint(note.CreatedBy, 10, 64)

	newNote := &models.Note{
		Title:     note.Title,
		Content:   note.Content,
		CreatedBy: uint(u64),
	}

	result := r.DB.Create(newNote)

	if result.Error != nil {
		return nil, result.Error
	}

	return &model.Note{
		ID:      fmt.Sprint(newNote.ID),
		Title:   newNote.Title,
		Content: newNote.Content,
	}, nil
}

func (r *mutationResolver) UserCreate(ctx context.Context, user model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
