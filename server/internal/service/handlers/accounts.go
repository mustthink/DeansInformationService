package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/types"
)

func (app *application) NewAuthorization(w http.ResponseWriter, r *http.Request) {
	app.enableCors(&w)
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var acc types.Account

	jsonErr := json.NewDecoder(r.Body).Decode(&acc)
	if jsonErr != nil {
		app.serverError(w, jsonErr)
	}

	res, err := app.data.NewAuth(&acc)
	if err != nil {
		app.serverError(w, err)
	}

	json_data, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		app.errorLog.Println(err)
	}

	fmt.Fprintf(w, "%v", string(json_data))
}
