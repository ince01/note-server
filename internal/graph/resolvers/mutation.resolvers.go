package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ince01/note-server/internal/auth"
	"github.com/ince01/note-server/internal/graph/generated"
	"github.com/ince01/note-server/internal/graph/model"
	"github.com/ince01/note-server/internal/orm/models"
	"github.com/ince01/note-server/pkg/helpers"
	"github.com/ince01/note-server/pkg/jwt"
)

func (r *mutationResolver) NoteCreate(ctx context.Context, note model.NoteInput) (*model.Note, error) {
	currentUser, _ := auth.ForContext(ctx)

	if note.Parent != nil {
		parentNote := &models.Note{}
		tx := r.DB.First(parentNote, *note.Parent)
		if tx.Error != nil {
			return nil, fmt.Errorf("parent note not found")
		}
	}

	parent := helpers.String2Uint(note.Parent)
	createdBy := helpers.String2Uint(&currentUser.ID)

	newNote := &models.Note{
		Title:     note.Title,
		Content:   note.Content,
		Icon:      note.Icon,
		Parent:    parent,
		CreatedBy: *createdBy,
	}

	result := r.DB.Create(newNote)
	if result.Error != nil {
		return nil, result.Error
	}

	return &model.Note{
		ID:        fmt.Sprint(newNote.ID),
		Title:     newNote.Title,
		Content:   newNote.Content,
		Parent:    note.Parent,
		CreatedBy: currentUser.ID,
	}, nil
}

func (r *mutationResolver) NoteUpdate(ctx context.Context, note model.NoteInput) (*model.Note, error) {
	updatedNote := models.Note{}

	tx := r.DB.First(&updatedNote, note.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	updatedNote.Title = note.Title
	updatedNote.Content = note.Content

	tx = r.DB.Save(&updatedNote)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model.Note{
		ID:        fmt.Sprint(updatedNote.ID),
		Title:     updatedNote.Title,
		Icon:      updatedNote.Icon,
		Content:   updatedNote.Content,
		CreatedBy: fmt.Sprint(updatedNote.CreatedBy),
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
		Gender:    string(user.Gender),
	}

	tx := r.DB.Where("email = ?", user.Email).First(newUser)

	if tx.RowsAffected > 0 {
		return nil, fmt.Errorf("user has been registered with this email")
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
		Gender:    model.Gender(newUser.Gender),
	}, nil
}

func (r *mutationResolver) TokenCreate(ctx context.Context, userCredential model.UserCredential) (*model.Token, error) {
	user := models.User{}

	tx := r.DB.Where("email = ?", userCredential.UserName).First(&user)

	if tx.Error != nil {
		return nil, fmt.Errorf("invaild email or password")
	}

	isMatchedPassword := user.ComparePassword(userCredential.Password)

	if !isMatchedPassword {
		return nil, fmt.Errorf("invaild email or password")
	}

	accessToken, err := jwt.GenerateToken(fmt.Sprint(user.ID))

	if err != nil {
		return nil, err
	}

	return &model.Token{
		TokenType:    model.TokenType(model.TokenTypeBearer),
		AccessToken:  accessToken.Jwt,
		ExpiresIn:    accessToken.Exp,
		RefreshToken: nil,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
