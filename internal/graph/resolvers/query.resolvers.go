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
)

func (r *queryResolver) Note(ctx context.Context, id int) (*model.Note, error) {
	currentUser, _ := auth.ForContext(ctx)

	note := models.Note{}

	tx := r.DB.
		Where(&models.Note{
			CreatedBy: *(helpers.String2Uint(&currentUser.ID)),
		}).
		First(&note, fmt.Sprint(id))

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model.Note{
		ID:        fmt.Sprint(note.ID),
		Title:     note.Title,
		Icon:      note.Icon,
		Content:   note.Content,
		Parent:    helpers.Uint2String(note.Parent),
		CreatedBy: fmt.Sprint(note.CreatedBy),
		CreatedAt: note.CreatedAt,
	}, nil
}

func (r *queryResolver) Notes(ctx context.Context, limit int, offset int) ([]model.Note, error) {
	currentUser, _ := auth.ForContext(ctx)

	var notes []models.Note

	tx := r.DB.
		Where(&models.Note{
			CreatedBy: *(helpers.String2Uint(&currentUser.ID)),
		}).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&notes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []model.Note

	for _, v := range notes {
		result = append(result, model.Note{
			ID:        fmt.Sprint(v.ID),
			Title:     v.Title,
			Icon:      v.Icon,
			Content:   v.Content,
			Parent:    helpers.Uint2String(v.Parent),
			CreatedAt: v.CreatedAt,
			CreatedBy: fmt.Sprint(v.CreatedBy),
		})
	}

	return result, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	currentUser, _ := auth.ForContext(ctx)

	var user models.User

	tx := r.DB.First(&user, currentUser.ID)

	if tx.Error != nil {
		return nil, tx.Error
	}

	result := &model.User{
		ID:        fmt.Sprint(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     &user.Phone,
		Gender:    model.Gender(user.Gender),
		AvatarURL: &user.AvatarUrl,
		CreatedAt: user.CreatedAt,
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
