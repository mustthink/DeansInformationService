package handlers

import "net/http"

func (app *application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/student/create", app.createStudent)

	return mux
}
