package api

import (
	"golangCourse/cloudsound/pkg/repository"
	songsservice "golangCourse/cloudsound/pkg/services/songservice"
	"net/http"
	"os"
	"log/slog"
	"github.com/gorilla/mux"
)

// Описание структуры API
type api struct {
	host string
	db 	 *repository.PGRepo
	r    *mux.Router
	logger *slog.Logger
}

// Функция-конструктор API
func New(host string, r *mux.Router, db *repository.PGRepo) *api {
	return &api{
		host: host, 
		r: r, 
		db: db, 
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)), 
	}
}

// Метод для запуска HTTP-сервера
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.host, api.r)
}

// Метод для конфигурации роутов
func (api *api) FillEndpoints() {
	api.r.Use(api.middleware)
	api.r.Handle("/", http.FileServer(http.Dir("../client")))
	api.r.HandleFunc("/api/songs", api.SongsHandler).Methods(
		http.MethodGet,
		http.MethodPost, 
		http.MethodPut,
	)
	api.r.HandleFunc("/api/songs/{id}", api.SongsHandler).Methods(
		http.MethodGet,
		http.MethodDelete,
	)
}

// Метод для передачи структуры api 
// в обработчик /api/songs
func (api *api) SongsHandler(w http.ResponseWriter, r *http.Request) {
	songsservice.ProcessSongsRequest(w, r, api.db)
}
