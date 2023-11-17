package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

// Функция для отправки записей переданного типа
// в формате JSON
func SendJsonRecords[T any](w http.ResponseWriter, records []T) {
	var jsonResponse []byte
	if len(records) == 1 {
		jsonResponse, err := json.Marshal(records[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
			} 
			w.Write(jsonResponse)
	}
	
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(records)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.Write(jsonResponse)
}
