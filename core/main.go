package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:3000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	infoLog.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB() (*sql.DB, error) {
	db, err := sql.Open(DB_STRING)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
