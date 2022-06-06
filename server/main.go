package main

import (
	"html/template"
	"log"
	"net/http"
	"server/types"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./html/temp.htm")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {

	var info = []string{"17", "March", "03", "1341", "Netudykhata", "Mykola", "Serhiiovych", "KBIKS-20-4", "2", "KIY", "Something very interesting", "31", "May", "22", "Lyashenko I.B.", "Petrenko A.S."}
	docx := types.СreateDoc(info)
	types.GenerateDoc("document", docx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	types.DeleteDoc()

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
