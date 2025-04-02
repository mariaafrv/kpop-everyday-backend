package main

import (
	"kpopeveryday/controller"
	"kpopeveryday/db"
	"kpopeveryday/repository"
	"kpopeveryday/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	MusicRepository := repository.NewMusicRepository(dbConnection)
	MusicUsecase := usecase.NewMusicUsecase(MusicRepository)
	MusicController := controller.NewMusicController(MusicUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pongg",
		})
	})

	server.GET("/music", MusicController.GetMusic)

	server.Run(":8000")
}
