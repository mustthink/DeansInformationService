package handlers

import (
	"fmt"
	"net/http"
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

	id, err := app.students.Insert(fio, keygroup, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/student?id=%d", id), http.StatusSeeOther)
}
