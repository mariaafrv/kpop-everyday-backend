package controller

import (
	"kpopeveryday/usecase"

	"github.com/gin-gonic/gin"
)

type MusicController struct {
	MusicUsecase usecase.MusicUsecase
}

func NewMusicController(usecase usecase.MusicUsecase) MusicController {
	return MusicController{
		MusicUsecase: usecase,
	}
}

func (p *MusicController) GetMusic(ctx *gin.Context) {

}
