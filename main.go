package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	_ = json.NewEncoder(writer).Encode(musics)

}

func getMusic(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "Application/json")
	par := mux.Vars(request)
	for _, music := range musics{
		if music.ID == par["id"] {
			_ = json.NewEncoder(writer).Encode(music)
			return
		}
	}

	_ = json.NewEncoder(writer).Encode(&Music{})
}

func createMusic(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "Application/json")
	var music Music
	_ = json.NewDecoder(request.Body).Decode(&music)
	music.ID = strconv.Itoa(rand.Intn(100000))
	musics = append(musics, music)
	_ = json.NewEncoder(writer).Encode(music)
}

func updateMusic(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "Application/json")
	par := mux.Vars(request)
	for index, music := range musics{
		if music.ID == par["id"] {
			musics = append(musics[:index], musics[index + 1:]...)
			var music Music
			_ = json.NewDecoder(request.Body).Decode(&music)
			music.ID = par["id"]
			musics = append(musics, music)
			_ = json.NewEncoder(writer).Encode(music)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(musics)
}

func deleteMusic(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "Application/json")
	par := mux.Vars(request)
	for index, music := range musics{
		if music.ID == par["id"] {
			musics = append(musics[:index], musics[index + 1:]...)
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(musics)
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
