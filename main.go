package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gameusecase "github.com/ricardoferrari/ginrest/usecases"
)

func main() {
	r := gin.Default()

	gameUC := gameusecase.NewGameUseCase()

	r.POST("/games", func(c *gin.Context) {
		var game gameusecase.Game
		if err := c.ShouldBindJSON(&game); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		gameUC.CreateGame(game)
		c.JSON(http.StatusCreated, game)
	})

	r.GET("/games/:id", func(c *gin.Context) {
		id := c.Param("id")
		game, err := gameUC.GetGame(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, game)
	})

	// List all games
	r.GET("/games", func(c *gin.Context) {
		c.JSON(http.StatusOK, gameUC.ListGames())
	})

	// Update a game
	r.PUT("/games/:id", func(c *gin.Context) {
		id := c.Param("id")
		var game gameusecase.Game
		if err := c.ShouldBindJSON(&game); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		game.ID = id

		err := gameUC.UpdateGame(game)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, game)
	})

	// Delete a game
	r.DELETE("/games/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := gameUC.DeleteGame(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
