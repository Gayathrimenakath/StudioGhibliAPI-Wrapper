package main

import (
	"log"
	"net/http"

	"github.com/StudioGhibliAPI-Wrapper/studiogbl"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("listening...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/movies", studiogbl.GetSpecies)
	log.Fatal(http.ListenAndServe(":8080", router))
}
