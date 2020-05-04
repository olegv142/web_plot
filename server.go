package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apiInit(r)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
