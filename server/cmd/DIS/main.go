package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"server/internal/db"
	"server/internal/service/handlers"
)

func main() {
	addr := flag.String("addr", "localhost:80", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "root:123456@tcp(127.0.0.1:3306)/dis_db?parseTime=true", "Название MySQL источника данных")
	url := flag.String("url", "http://localhost:8080", "URL Web сайта")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	secret := "secret"

	db, err := db.OpenDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := handlers.NewApplication(errorLog, infoLog, db, url, secret)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	app.InfoLog().Printf("Запуск веб-сервера на %s", *addr)
	err = srv.ListenAndServe()
	app.ErrorLog().Fatal(err)
}
