package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

// Функция для отправки записей переданного типа
// в формате JSON
func SendJsonRecords[T any](w http.ResponseWriter, records []T) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(records)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonResponse)
}
