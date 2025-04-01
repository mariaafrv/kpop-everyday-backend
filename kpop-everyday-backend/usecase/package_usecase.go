package usecase

import "kpopeveryday/model"

type MusicUsecase struct {
}

func NewMusicUsecase() MusicUsecase {
	return MusicUsecase{}
}

func (pu *MusicUsecase) GetMusic() ([]model.Music, error) {
	return []model.Music, nil
}
