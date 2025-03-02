package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoferrari/ginrest/usecases/gameusecase"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
