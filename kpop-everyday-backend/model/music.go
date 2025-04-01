package model

import "time"

type Music struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Artist  string    `json:"artist"`
	Link    string    `json:"youtube_link"`
	Release time.Time `json:"release_date"`
}
