package config

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rif123/go-react/controllers"
	"net/http"
)

func SetRouter () *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))


	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	api.GET("/jokes", controllers.JokeHandler)
	api.POST("/jokes/like/:jokeID", controllers.LikeJoke)

	return router
}