package main

import (
	"encoding/json"
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
	Album   string `json:"record"`
}

var musics []Music

func getMusics(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(writer).Encode(musics)

}

func getMusic(writer http.ResponseWriter, request *http.Request)  {

}

func createMusic(writer http.ResponseWriter, request *http.Request)  {

}

func updateMusic(writer http.ResponseWriter, request *http.Request)  {

}

func deleteMusic(writer http.ResponseWriter, request *http.Request)  {

}

func main()  {
	router := mux.NewRouter()

	//Samples
	musics = append(musics, Music{ID: "1", Title: "My Immortal", Singer: &Singer{Firstname: "Evanescence", Lastname: "", Album: "Origin"}})
	musics = append(musics, Music{ID: "2", Title: "Lithium", Singer: &Singer{Firstname: "Evanescence", Lastname: "", Album: "The Open Door"}})
	musics = append(musics, Music{ID: "3", Title: "Bring Me to Life", Singer: &Singer{Firstname: "Evanescence", Lastname: "", Album: "Fallen"}})

	//Endpoints
	router.HandleFunc("/api/musics", getMusics).Methods("GET")
	router.HandleFunc("/api/musics/{id}", getMusic).Methods("GET")
	router.HandleFunc("/api/musics", createMusic).Methods("POST")
	router.HandleFunc("/api/musics/{id}", updateMusic).Methods("PUT")
	router.HandleFunc("/api/musics/{id}", deleteMusic).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
