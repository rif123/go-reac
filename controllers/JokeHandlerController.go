package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rif123/go-react/models"
	"net/http"
)

func JokeHandler(c *gin.Context) {

	jokes := []models.Joke{
		{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
		{2, 0, "What do you call a fake noodle? An Impasta."},
		{3, 0, "How many apples grow on a tree? All of them."},
		{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
		{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
		{6, 0, "Why did the coffee file a police report? It got mugged."},
		{7, 0, "How does a penguin build it's house? Igloos it together."},
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

