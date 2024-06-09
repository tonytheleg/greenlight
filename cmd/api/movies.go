package main

import (
	"fmt"
	"net/http"
)

// creatMovieHandler creates movies
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler gets movies, any interpolated URL params will be stored
// in the request context, we use ParamsFromContext to get those values
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		http.NotFound(w, r)
		return

	}

	fmt.Fprintf(w, "Show the details of movie %d\n", id)
}
