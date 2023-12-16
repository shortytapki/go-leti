package songsservice

import (
	// "encoding/json"
	// "fmt"
	// "golangCourse/cloudsound/pkg/entities"
	// "golangCourse/cloudsound/pkg/lib"

	// "log"
	"encoding/json"
	"fmt"
	"golangCourse/cloudsound/pkg/entities"
	"golangCourse/cloudsound/pkg/lib"
	"golangCourse/cloudsound/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "sync"
)

// Функция обработки эндпоинта /api/songs
func ProcessSongsRequest(w http.ResponseWriter, r *http.Request, db *repository.PGRepo) {
	// var mutex sync.Mutex
	switch r.Method {
	case http.MethodGet:
		params := mux.Vars(r)
		id, ok := params["id"]
		if ok {
			id, _ := strconv.Atoi(id)
			song, err := db.GetSongById(id) 
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 Internal "))
				return 
			}
			lib.SendJsonRecords(w, []entities.Song{song})
			return
		}
		songs, err := db.GetSongs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%v Internal server error.", http.StatusInternalServerError)))
			return 
		}
		lib.SendJsonRecords(w, songs)
		
	case http.MethodPost:
		var newSong entities.Song
		err := json.NewDecoder(r.Body).Decode(&newSong)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%v Bad Request", http.StatusBadRequest)))
			return
		}
		err = db.AddSong(newSong)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Cloudsound database is temporarily down, try again later."))
			return
		}
		w.Write([]byte(fmt.Sprintf("The song %v is successfully added.", newSong.Name)))
		
	case http.MethodPut:
		var newSong entities.Song
		err := json.NewDecoder(r.Body).Decode(&newSong)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%v Bad Request", http.StatusBadRequest)))
			return
		}
		db.UpdateSong(&newSong)
		w.Write([]byte(fmt.Sprintf("The song %v is successfully updated.", newSong.Name)))
		
	case http.MethodDelete:
		params := mux.Vars(r)
		id, ok := params["id"]
		if ok {
			// Тут не хватает логики проверки
			// существования песни с переданным id!
			id, _ := strconv.Atoi(id)
			db.DeleteSong(id)
			w.Write([]byte(fmt.Sprintf("The song with id:%v is successfully deleted.", id)))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v Bad Request", http.StatusBadRequest)))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
