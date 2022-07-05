package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/service/types"
	"strconv"
)

func (app *application) createNews(w http.ResponseWriter, r *http.Request) {
	app.enableCors(&w)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	var nw types.News
	jsonErr := json.NewDecoder(r.Body).Decode(&nw)

	if jsonErr != nil {
		app.serverError(w, jsonErr)
	}

	err := app.data.InsertNews(&nw)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {
	app.enableCors(&w)
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.data.GetNews(id)
	if err != nil {
		if errors.Is(err, types.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	json_data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%v", string(json_data))
}

func (app *application) showListNews(w http.ResponseWriter, r *http.Request) {
	app.enableCors(&w)
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	s, err := app.data.LatestNews()

	if err != nil {
		app.serverError(w, err)
		return
	}

	json_data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		app.serverError(w, err)
	}

	fmt.Fprintf(w, "%v", string(json_data))
}
