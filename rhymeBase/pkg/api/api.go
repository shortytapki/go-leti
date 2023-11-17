package api

import (
	"golangCourse/rhymeBase/pkg/services/database"
	"golangCourse/rhymeBase/pkg/services/rhymesService"
	songsservice "golangCourse/rhymeBase/pkg/services/songsService"
	"net/http"
)

// Описание структуры API
type api struct {
	DB   *database.DB
	host string
	r    *http.ServeMux
}

// Функция-конструктор API
func New(host string, r *http.ServeMux) *api {
	return &api{host: host, r: r, DB: database.New()}
}

// Метод для запуска HTTP-сервера
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.host, api.r)
}

// Метод для конфигурации роутов
func (api *api) FillEndpoints() {
	api.r.Handle("/", http.FileServer(http.Dir("../client")))
	api.r.HandleFunc("/api/rhymes", api.RhymesHandler)
	api.r.HandleFunc("/api/songs", api.SongsHandler)
}

// Метод для передачи структуры api 
// в обработчик /api/rhymes
func (api *api) RhymesHandler(w http.ResponseWriter, r *http.Request) {
	rhymesService.ProcessRhymeRequest(w, r, api.DB)
}

// Метод для передачи структуры api 
// в обработчик /api/songs
func (api *api) SongsHandler(w http.ResponseWriter, r *http.Request) {
	songsservice.ProcessSongsRequest(w, r, api.DB)
}
