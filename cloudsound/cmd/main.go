package main

import (
	"fmt"
	"golangCourse/cloudsound/pkg/api"
	"golangCourse/cloudsound/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)
const HOST = "127.0.0.1:8080"
const ENV_PATH = "../../"
func main() {
	db, err := repository.New(ENV_PATH)
	if err != nil {
		log.Fatal("Could not connect to the database.")
	}
	api := api.New(HOST, mux.NewRouter(), db)
	api.FillEndpoints()
	fmt.Printf("The server is started on http://%v\n", HOST)
	log.Fatal(api.ListenAndServe())
}