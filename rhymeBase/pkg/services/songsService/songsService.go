package songsservice

import (
	"encoding/json"
	"fmt"
	"golangCourse/rhymeBase/pkg/entities"
	"golangCourse/rhymeBase/pkg/lib"
	"golangCourse/rhymeBase/pkg/services/database"
	"log"
	"net/http"
	"sync"
)

// Функция обработки эндпоинта /api/songs
func ProcessSongsRequest(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var mutex sync.Mutex
	switch r.Method {
	case http.MethodGet:
		lib.SendJsonRecords(w, db.Rhymes)
		return
	case http.MethodPost:
		var newSong entities.Song
		err := json.NewDecoder(r.Body).Decode(&newSong)
		if err != nil {
			log.Println(err)
		}
		// Защита записи
		mutex.Lock()
		db.Songs = append(db.Songs, entities.Song{Title: newSong.Title, Artist: newSong.Artist})
		mutex.Unlock()
		w.Write([]byte(fmt.Sprintf("The song %v by %v is added.\n", newSong.Title, newSong.Artist)))
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
