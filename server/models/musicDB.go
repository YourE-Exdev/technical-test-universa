package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Music, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, album, albumart, singer, 
				publishdate, createat, updateat   
				FROM 
					music 
				WHERE id =$1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var music Music

	err := row.Scan(
		&music.ID,
		&music.Name,
		&music.Album,
		&music.AlbumArt,
		&music.Singer,
		&music.PublishDate,
		&music.CreateAt,
		&music.UpdateAt,
	)

	if err != nil {
		return nil, err
	}

	return &music, nil
}

func (m *DBModel) All() ([]*Music, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name, album, albumart, singer, publishdate, createat, updateat FROM music`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musics []*Music

	for rows.Next() {
		var music Music
		err := rows.Scan(
			&music.ID,
			&music.Name,
			&music.Album,
			&music.AlbumArt,
			&music.Singer,
			&music.PublishDate,
			&music.CreateAt,
			&music.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		musics = append(musics, &music)
	}
	return musics, nil
}

func (m *DBModel) InsertMusic(music Music) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO music (name, album, albumart, singer, publishdate, createat, updateat) 
				values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, query,
		music.Name,
		music.Album,
		music.AlbumArt,
		music.Singer,
		music.PublishDate,
		music.CreateAt,
		music.UpdateAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateMusic(music Music) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE music set name = $1, album = $2, albumart=$3 , singer=$4 , publishdate=$5, updateat=$6
				Where id=$7`

	_, err := m.DB.ExecContext(ctx, query,
		music.Name,
		music.Album,
		music.AlbumArt,
		music.Singer,
		music.PublishDate,
		music.UpdateAt,
		music.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteMusic(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM music WHERE id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
