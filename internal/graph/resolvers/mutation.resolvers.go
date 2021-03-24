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
		ID:        fmt.Sprint(newNote.ID),
		Title:     newNote.Title,
		Content:   newNote.Content,
		CreatedBy: note.CreatedBy,
	}, nil
}

func (r *mutationResolver) UserCreate(ctx context.Context, user model.UserInput) (*model.User, error) {
	newUser := &models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		AvatarUrl: user.AvatarURL,
	}

	tx := r.DB.Where("email = ?", user.Email).First(newUser)

	if tx.RowsAffected > 0 {
		return nil, fmt.Errorf("User has been registered with this email.")
	}

	tx = r.DB.Create(newUser)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model.User{
		ID:        fmt.Sprint(newUser.ID),
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     &newUser.Phone,
		AvatarURL: &newUser.AvatarUrl,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

func (r *mutationResolver) UserDelete(ctx context.Context, id string) (*model.User, error) {
	var user models.User

	tx := r.DB.First(&user, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = r.DB.Delete(&user)

	if tx.Error != nil {
		return nil, tx.Error
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
