package api

import (
	"encoding/json"
	"fmt"
	"golangCourse/rhymeBase/pkg/servicedb"
	"log"
	"net/http"
)

// Описание структуры API
type api struct {
	host string
	r    *http.ServeMux
	db   *servicedb.DB
}

// Функция-конструктор API
func New(host string, r *http.ServeMux) *api {
	rhymes := make([]servicedb.Rhyme, 0)
	songs := make([]servicedb.Song, 0)
	return &api{host: host, r: r, db: servicedb.New(rhymes, songs)}
}

// Метод для запуска HTTP-сервера
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.host, api.r)
}

// Метод для конфигурации роутов
func (api *api) FillEndpoints() {
	api.r.Handle("/", http.FileServer(http.Dir("../client")))
	api.r.HandleFunc("/api/rhymes", RhymesHandler(api.db))
	api.r.HandleFunc("/api/songs", SongsHandler(api.db))
}

// Функция для отправки записей переданного типа 
// в формате JSON
func sendJsonRecords[T servicedb.Rhyme | servicedb.Song](w http.ResponseWriter, records []T) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(records)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResponse)	
}

// Функция обработки эндпоинта /api/rhymes
// аргументом принимает базу для передачи в обработчик
func RhymesHandler(db *servicedb.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			sendJsonRecords(w, db.Rhymes)
			return
		case http.MethodPost:
			var newRhyme servicedb.Rhyme
			err := json.NewDecoder(r.Body).Decode(&newRhyme)
			if err != nil {
				log.Fatal(err)
			}
			db.Rhymes = append(db.Rhymes, servicedb.Rhyme{Text: newRhyme.Text})
			w.Write([]byte(fmt.Sprintf("%v rhyme is added.\n", newRhyme.Text)))
			return
		default:
			w.Write([]byte("Unknown request method."))
		}
	}
	
}

// Функция обработки эндпоинта /api/songs
// аргументом принимает базу для передачи в обработчик
func SongsHandler(db *servicedb.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			sendJsonRecords(w, db.Songs)
			return
		case http.MethodPost:
			var newSong servicedb.Song
			err := json.NewDecoder(r.Body).Decode(&newSong)
			if err != nil {
				log.Fatal(err)
			}
			db.Songs = append(db.Songs, servicedb.Song{Title: newSong.Title,Artist: newSong.Artist})
			w.Write([]byte(fmt.Sprintf("The song %v by %v is added.\n", newSong.Title, newSong.Artist)))
			return
		default:
			w.Write([]byte("Unknown request method."))
		}
	}

}
