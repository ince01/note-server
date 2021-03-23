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

func (r *queryResolver) Note(ctx context.Context, id int) (*model.Note, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Notes(ctx context.Context, limit *int, offset *int) ([]model.Note, error) {
	var notes []models.Note

	r.DB.Limit(*limit).Offset(*offset).Find(&notes)

	// err := r.DB.Model(&models.Note{}).Association("CreatedBy").Find(&notes)

	var result []model.Note

	for _, v := range notes {
		result = append(result, model.Note{
			ID:        fmt.Sprint(v.ID),
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			CreatedBy: &model.User{
				ID: fmt.Sprint(v.CreatedBy),
			},
		})
	}

	return result, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
