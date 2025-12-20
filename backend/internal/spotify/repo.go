package iSpotify

import (
	"database/sql"
)

type SpotifyRepo struct {
	db *sql.DB
}

func NewSpotifyRepo(db *sql.DB) *SpotifyRepo {
	return &SpotifyRepo{db}
}
