package controller

import (
	"kpopeveryday/usecase"
	"net/http"

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
	music, err := p.MusicUsecase.GetMusic()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, music)
}
