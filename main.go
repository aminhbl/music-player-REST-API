package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Music struct {
	ID   string `json:"id"`
	Title   string `json:"title"`
	Singer   *Singer `json:"singer"`
}

type Singer struct {
	Firstname   string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Record   string `json:"record"`
}

func main()  {
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/api/musics", getMusics).Methods("GET")
	router.HandleFunc("/api/musics/{id}", getMusic).Methods("GET")
	router.HandleFunc("/api/musics", createMusic).Methods("POST")
	router.HandleFunc("/api/musics/{id}", updateMusic).Methods("PUT")
	router.HandleFunc("/api/musics/{id}", deleteMusic).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
