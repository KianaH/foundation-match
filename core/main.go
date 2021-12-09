package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
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

	r := routes

	infoLog.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe("127.0.0.1:4000", r)
	errorLog.Fatal(err)

}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv(DB_URI))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
