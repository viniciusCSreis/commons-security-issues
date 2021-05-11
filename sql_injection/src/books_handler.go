package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type BooksHandler struct {
	PgConn PostgresConnector
}

func NewBooksHandler(PgConn PostgresConnector) BooksHandler {
	return BooksHandler{
		PgConn: PgConn,
	}
}

func (h BooksHandler) GET() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")

		dbConn := h.PgConn.DBConnection()

		SQL := fmt.Sprintf(`SELECT id, name FROM book WHERE name ilike ('%s')`, "%"+name+"%")


		log.Printf("SQL:%s", SQL)

		result, err := dbConn.Query(SQL)
		if err != nil {
			panic(err)
		}

		books := make([]Book, 0)

		for result.Next() {
			book := Book{}
			err := result.Scan(&book.Id, &book.Name)
			if err != nil {
				panic(err)
			}
			books = append(books, book)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		err = json.NewEncoder(writer).Encode(books)
		if err != nil {
			panic(err)
		}

	})
}
