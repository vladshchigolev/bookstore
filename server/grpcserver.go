package server

import (
	"bookstore/pkg/api"
	"context"
	"database/sql"
	"fmt"
)

type GRPCServer struct {
	DB *sql.DB
}

func New() *GRPCServer {
	return &GRPCServer{}
}
func (s *GRPCServer) ConfigureDatabase() error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *GRPCServer) GetBooks(ctx context.Context, auth *api.Author) (*api.BooksSet, error) {
	// здесь мы должны взять значение поля Name переданного методу объекта типа Author и по этому значению искать в базе книги
	//brepo := store.BookRepository{}
	//return nil, nil

	rows, err := s.DB.Query("SELECT isbn, title, author, year FROM books WHERE author = ?", auth.Name)
	fmt.Println(auth.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks api.BooksSet

	for rows.Next() {
		var bk api.Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Year)
		if err != nil {
			return nil, err
		}

		bks.Books = append(bks.Books, &bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &bks, nil
}

func (s *GRPCServer) GetAuthors(ctx context.Context, book *api.Title) (*api.AuthorsSet, error) {
	return nil, nil
}
