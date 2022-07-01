package main

import (
	"encoding/json" // encode and decode json
	"fmt"
	"log" // log errors
	"math/rand"
	"net/http"    // create server in golang
	"strconv"
	"github.com/gorilla/mux" // create router in golang
)

type Movie struct {
	ID string `json:"id"`
	Rating int `json:"rating"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var movies []Movie	

func main(){
	movies = append(movies, Movie{ID: "1", Rating: 3, Title: "The Matrix", Director: &Director{FirstName: "joe", LastName: "smith"}})
	movies = append(movies, Movie{ID: "4", Rating: 7, Title: "The MAtrix Reloaded", Director: &Director{FirstName: "lmfao", LastName: "lmao"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", createMovie).Methods("POST")

	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}	
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)

	for _, item := range movies{ // blank identifier is used to ignore the index
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	//object := Movie{ID: params["ID"], Rating: params["Rating"], Title: params["Title"], Director: &Director{FirstName: params["Director_fname"], LastName: params["Director_lname"]}}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if params["ID"] == item.ID {
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
		json.NewEncoder(w).Encode(movies)
	}
	//object := Movie{ID: params["ID"], Rating: params["Rating"], Title: params["Title"], Director: &Director{FirstName: params["Director_fname"], LastName: params["Director_lname"]}}
	//movies = append(movies, object)
	
	//create string list
}