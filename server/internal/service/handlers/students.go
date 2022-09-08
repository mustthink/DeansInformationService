package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"server/internal/types"
	"strconv"
)

func (app *application) createStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	keygroup := 1
	fio := "Netudykhata Mykola Serhiyovych"
	expires := 155

	err := app.data.InsertStudent(fio, keygroup, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) showStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.data.GetStudent(id)
	if err != nil {
		if errors.Is(err, types.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", s)
}

func (app *application) showListStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	s, err := app.data.LatestStudents()

	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, student := range s {
		fmt.Fprintf(w, "%v\n", student)
	}
}
