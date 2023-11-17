package database

import (
	"golangCourse/rhymeBase/pkg/entities"
)

// Описание структуры базы данных
type DB struct {
	Rhymes []entities.Rhyme
	Songs  []entities.Song
}

// Создаёт и возвращает указатель
// на структуру базы данных
func New() *DB {
	rhymes := make([]entities.Rhyme, 0)
	songs := make([]entities.Song, 0)
	return &DB{Rhymes: rhymes, Songs: songs}
}
