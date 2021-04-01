package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ince01/note-server/internal/graph/model"
	"github.com/ince01/note-server/internal/orm/models"
	"github.com/ince01/note-server/pkg/jwt"
	"gorm.io/gorm"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		if tokenStr == "" {
			c.Next()
			return
		}

		//validate jwt token
		userId, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.Error(err)
			return
		}

		user := models.User{}

		db.First(&user, userId)

		ctx := context.WithValue(c.Request.Context(), userCtxKey, model.User{
			ID:        fmt.Sprint(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     &user.Phone,
			AvatarURL: &user.AvatarUrl,
			CreatedAt: user.CreatedAt,
		})

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*model.User, error) {
	userContext := ctx.Value(userCtxKey)

	if userContext == nil {
		return nil, fmt.Errorf("could not retrieve gin.Context")
	}

	user, ok := userContext.(model.User)

	if !ok {
		return nil, fmt.Errorf("userContext has wrong type")
	}

	return &user, nil
}
