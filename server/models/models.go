package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Music struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Album       string    `json:"album"`
	AlbumArt    string    `json:"albumart"`
	Singer      string    `json:"singer"`
	PublishDate string    `json:"publishdate"`
	CreateAt    time.Time `json:"createat"`
	UpdateAt    time.Time `json:"updateat"`
}
