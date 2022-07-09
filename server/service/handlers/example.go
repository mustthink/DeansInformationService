package handlers

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	//var info = []string{"17", "March", "03", "1341", "Netudykhata", "Mykola", "Serhiiovych", "KBIKS-20-4", "2", "KIY", "Something very interesting", "31", "May", "22", "Lyashenko I.B.", "Petrenko A.S."}
	//docx := types.СreateDoc(info)
	//types.GenerateDoc("document", docx)
	//defer types.DeleteDoc()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ts, err := template.ParseFiles("./html/document.html")
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
}
