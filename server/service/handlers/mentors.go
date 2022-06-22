package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"server/service/types"
	"strconv"
)

func (app *application) createMentor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	mntr := &types.Mentor{ID: 0, FIO: "Vyohin He Ll", CHAIR: "CYBERSECURE"}

	err := app.data.InsertMentor(mntr)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showMentor(w http.ResponseWriter, r *http.Request) {
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

	s, err := app.data.GetMentor(id)
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
