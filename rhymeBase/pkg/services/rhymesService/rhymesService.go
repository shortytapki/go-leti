package rhymesService

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

// Функция обработки эндпоинта /api/rhymes
func ProcessRhymeRequest(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var mutex sync.Mutex
	switch r.Method {
	case http.MethodGet:
		lib.SendJsonRecords(w, db.Rhymes)
		return
	case http.MethodPost:
		var newRhyme entities.Rhyme
		err := json.NewDecoder(r.Body).Decode(&newRhyme)
		if err != nil {
			log.Println(err)
		}
		// Защита записи
		mutex.Lock()
		db.Rhymes = append(db.Rhymes, entities.Rhyme{Text: newRhyme.Text})
		mutex.Unlock()
		w.Write([]byte(fmt.Sprintf("%v rhyme is added.\n", newRhyme.Text)))
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
