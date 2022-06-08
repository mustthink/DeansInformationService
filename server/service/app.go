package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"server/service/handlers"
)

type App interface {
	home(w http.ResponseWriter, r *http.Request)
	Routes() *http.ServeMux
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	students *handlers.StudentModel
}

func NewApplication(errorLog *log.Logger, infoLog *log.Logger, db *sql.DB) *application {
	return &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		students: &handlers.StudentModel{DB: db},
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
