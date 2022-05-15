package main

import (
	"bookstore/models"
	"bookstore/pkg/api"
	"bookstore/server"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	var err error

	// Initalize the sql.DB connection pool and assign it to the models.DB
	// global variable.
	models.DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := server.GRPCServer{}
	srv.ConfigureDatabase()
	api.RegisterBookStorageServer(s, &srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
	//http.HandleFunc("/books", booksIndex)
	//http.ListenAndServe(":8080", nil)
}

// booksIndex sends a HTTP response listing all books.
func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Println(bk.Isbn, bk.Title, bk.Author, bk.Year)
		//fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Year)
	}
}
