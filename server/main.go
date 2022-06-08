package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/service/handlers"
	"server/types"
)

func main() {
	addr := flag.String("addr", "localhost:8081", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "root:123456@/dis_db?parseTime=true", "Название MySQL источника данных")
	flag.Parse()

	var info = []string{"17", "March", "03", "1341", "Netudykhata", "Mykola", "Serhiiovych", "KBIKS-20-4", "2", "KIY", "Something very interesting", "31", "May", "22", "Lyashenko I.B.", "Petrenko A.S."}
	docx := types.СreateDoc(info)
	types.GenerateDoc("document", docx)
	defer types.DeleteDoc()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := db.OpenDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := handlers.NewApplication(errorLog, infoLog, db)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	app.InfoLog().Printf("Запуск веб-сервера на %s", *addr)
	err = srv.ListenAndServe()
	app.ErrorLog().Fatal(err)
}
