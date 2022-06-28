package handlers

import "net/http"

func (app *application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/student", app.showStudent)
	mux.HandleFunc("/student/create", app.createStudent)
	mux.HandleFunc("/student/list", app.showListStudents)
	mux.HandleFunc("/mentor", app.showMentor)
	mux.HandleFunc("/mentor/create", app.createMentor)
	mux.HandleFunc("/news", app.showNews)
	mux.HandleFunc("/news/create", app.createNews)
	mux.HandleFunc("/news/list", app.showListNews)

	return mux
}
