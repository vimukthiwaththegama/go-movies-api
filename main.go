package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"stringconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie
var directors []Director

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/movies",getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovieById).Methods("GET")
	r.HandleFunc("movies",saveMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting the application on port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}

func getAllMovies(){

}