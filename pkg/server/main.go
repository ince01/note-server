package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ince01/note-server/internal/graph"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// Run the server
func Run(db *gorm.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GraphQL Routes
	r.POST("/graphql", graph.Handler(db))
	r.GET("/graphql", graph.PlaygroundHandler())

	r.Run(":" + defaultPort)
}
