package main

import (
	"fmt"
	"github.com/viniciuscsreis/commons-security-issues/sql_injection/src"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Hello World\n")

	dbConfig := src.DataBaseConfiguration{
		Type:            "pgx",
		Host:            "localhost",
		Port:            5432,
		Username:        "sql_injection_poc",
		Password:        "sql_injection_poc",
		DBName:          "sql_injection_poc",
		SSL:             "disable",
		ConnMaxLifetime: 600,
		MaxIdleConns:    25,
		MaxOpenConns:    25,
	}
	pgConnector := src.NewPostgresConnector(dbConfig)

	booksHandler := src.NewBooksHandler(pgConnector)

	http.HandleFunc("/books", booksHandler.GET().ServeHTTP)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
