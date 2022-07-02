package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"server/service/data"
)

type App interface {
	home(w http.ResponseWriter, r *http.Request)
	Routes() *http.ServeMux
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	data     *data.Service
	url      *string
}

func NewApplication(errorLog *log.Logger, infoLog *log.Logger, db *sql.DB, url *string) *application {
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		data:     &data.Service{DB: db},
		url:      url,
	}
}

func (a *application) ErrorLog() *log.Logger {
	return a.errorLog
}

func (a *application) InfoLog() *log.Logger {
	return a.infoLog
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", *app.url)
}
