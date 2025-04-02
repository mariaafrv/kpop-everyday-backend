package usecase

import (
	"kpopeveryday/model"
	"kpopeveryday/repository"
)

type MusicUsecase struct {
	repository repository.MusicRepository
}

func NewMusicUsecase(repo repository.MusicRepository) MusicUsecase {
	return MusicUsecase{
		repository: repo,
	}
}

func (pu *MusicUsecase) GetMusic() ([]model.Music, error) {
	return pu.repository.GetMusic()
}
