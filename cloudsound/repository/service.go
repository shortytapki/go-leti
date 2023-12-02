package repository

import (
	"context"
	"golangCourse/cloudsound/entities"
)
// Функция для создания записи 
// песни в БД 
func (repo *PGRepo) AddSong(song entities.Song) (error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	_, err := repo.pool.Exec(context.Background(), `INSERT INTO songs (name, duration, genre_id, artist_id) VALUES ($1, $2, $3, $4);`,
		song.Name,
		song.Duration,
		song.GenreId,
		song.ArtistId,
	)
	if err != nil {
		return err
	}
	return nil
}

// Функция для получения всех
// песен из БД
func (repo *PGRepo) GetSongs() ([]entities.Song, error){
	rows, err := repo.pool.Query(context.Background(), 
	`SELECT id, name, duration, genre_id, artist_id FROM songs;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var data []entities.Song
	for rows.Next() {
		var song entities.Song
		rows.Scan(
			&song.Id,
			&song.Name,
			&song.Duration,
			&song.GenreId,
			&song.ArtistId,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, song)
	}

	return data, nil
}

// Функция полученя песни
// по переданному id
func (repo *PGRepo) GetSongById(id int) (entities.Song, error) {
	row := repo.pool.QueryRow(
		context.Background(), 
		`SELECT id, name, duration, genre_id, artist_id FROM songs WHERE id=$1`,
		id,
	)
	var song entities.Song
	row.Scan(&song.Id, &song.Name, &song.Duration, &song.GenreId, &song.ArtistId)
	return song, nil
}

// Функция для изменения
// записи песни в БД
func (repo *PGRepo) UpdateSong(song *entities.Song) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.pool.QueryRow(context.Background(), `UPDATE songs SET name=$1, duration=$2 WHERE id=$3`, &song.Name, &song.Duration, &song.Id)
}

// Функция для удаления записи 
// в БД
func (repo *PGRepo) DeleteSong(id int) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.pool.QueryRow(context.Background(), `DELETE FROM songs WHERE id=$1`, id)
}