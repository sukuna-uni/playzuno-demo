package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie represents a movie model
type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     string `json:"year"`
}

// In-memory movie list (simulating a database)
var movies []Movie

// Get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get a single movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

// Create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}

// Delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Mock data
	movies = append(movies, Movie{ID: "1", Title: "Inception", Director: "Christopher Nolan", Year: "2010"})
	movies = append(movies, Movie{ID: "2", Title: "Interstellar", Director: "Christopher Nolan", Year: "2014"})
	movies = append(movies, Movie{ID: "3", Title: "GOAT", Director: "Makatha", Year: "2014"})

	// Define the routes
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
