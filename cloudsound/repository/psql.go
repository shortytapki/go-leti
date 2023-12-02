package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

// Описание структуры пула соединений
type PGRepo struct {
	mutex sync.Mutex
	pool *pgxpool.Pool
}

// Функция создания нового репозитория PostgreSQL
func New (relativeEnvPath string) (*PGRepo, error) {
	path, err := filepath.Abs("")
	if err != nil {
		return &PGRepo{}, err
	}
	err = godotenv.Load(filepath.Join(path, relativeEnvPath, ".env"))

	if err != nil {
		return &PGRepo{}, err
	}
	connStr := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v", 
		os.Getenv("DB_USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DB"))

	pool, err := pgxpool.Connect(context.Background(), connStr)
	
	if err != nil {
		return nil, err
	}
	return &PGRepo{mutex: sync.Mutex{}, pool: pool}, nil
}