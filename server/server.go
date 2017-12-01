package server

import (
	"net/http"

	"github.com/sulavvr/neogo/controllers"
)

var (
	MovieController controllers.Movie
	GenreController controllers.Genre
)

/**
 * Start handles the routes and also starts the server to serve the application.
 */
func Start() {
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// http.HandleFunc("/", MovieController.Index)
	http.HandleFunc("/movies/", MovieController.Index)
	http.HandleFunc("/movie/", MovieController.Show)
	http.HandleFunc("/genres/", GenreController.Index)
	http.HandleFunc("/genre/", GenreController.Show)

	http.ListenAndServe(":3000", nil)
}
