package main

import (
	"fmt"
	"golangCourse/rhymeBase/pkg/api"
	"log"
	"net/http"
)

const HOST = "127.0.0.1:8090"

func main() {
	api := api.New(HOST, http.NewServeMux())
	api.FillEndpoints()
	fmt.Printf("The server is started on http://%v\n", HOST)
	log.Fatal(api.ListenAndServe())
}
