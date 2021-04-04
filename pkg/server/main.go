package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ince01/note-server/internal/auth"
	"github.com/ince01/note-server/internal/graph"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// Run the server
func Run(db *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	router.Use(cors.AllowAll())

	router.Use(auth.Middleware(db))

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GraphQL Routes
	router.POST("/graphql", graph.Handler(db))
	router.GET("/graphql", graph.PlaygroundHandler())

	router.Run(":" + defaultPort)
}
