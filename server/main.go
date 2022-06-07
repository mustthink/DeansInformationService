package main

import (
	"flag"
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
	addr := flag.String("addr", "localhost:8081", "Сетевой адрес HTTP")
	flag.Parse()

	var info = []string{"17", "March", "03", "1341", "Netudykhata", "Mykola", "Serhiiovych", "KBIKS-20-4", "2", "KIY", "Something very interesting", "31", "May", "22", "Lyashenko I.B.", "Petrenko A.S."}
	docx := types.СreateDoc(info)
	types.GenerateDoc("document", docx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Запуск веб-сервера на %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
