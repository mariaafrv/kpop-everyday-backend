package repository

import (
	"database/sql"
	"fmt"
	"kpopeveryday/model"
)

type MusicRepository struct {
	connection *sql.DB
}

func NewMusicRepository(connection *sql.DB) MusicRepository {
	return MusicRepository{
		connection: connection,
	}
}

func (pr *MusicRepository) GetMusic() ([]model.Music, error) {
	query := "SELECT id, name, artist, youtube_link, release_date FROM music"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Music{}, err
	}
	var musicList []model.Music
	var musicObj model.Music

	for rows.Next() {
		err = rows.Scan(
			&musicObj.ID,
			&musicObj.Name,
			&musicObj.Artist,
			&musicObj.Link,
			&musicObj.Release,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Music{}, err
		}

		musicList = append(musicList, musicObj)
	}

	rows.Close()
	return musicList, nil
}
