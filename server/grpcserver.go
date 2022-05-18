package server

import (
	"bookstore/pkg/api"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type GRPCServer struct {
	config *Config
	DB     *sql.DB
	logger *logrus.Logger
}

type Config struct {
	DataSourceName string `toml:"data_source_name"`
	LogLevel       string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{}
}
func New(config *Config) *GRPCServer {
	return &GRPCServer{config: config, logger: logrus.New()}
}
func (s *GRPCServer) ConfigureDatabase() error {
	db, err := sql.Open("mysql", s.config.DataSourceName)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.logger.Info("db connection established")
	s.DB = db
	return nil
}
func (s *GRPCServer) ConfigureLogger() error {
	//logger := logrus.New()
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.Info("logger configured")
	return nil
}
func (s *GRPCServer) GetBooks(ctx context.Context, auth *api.Author) (*api.BooksSet, error) {
	// здесь мы должны взять значение поля Name переданного методу объекта типа Author и по этому значению искать в базе книги
	//brepo := store.BookRepository{}
	//return nil, nil

	rows, err := s.DB.Query("SELECT books.isbn, title, year FROM books, authors WHERE author = ? AND books.isbn = authors.isbn", auth.Name)
	s.logger.Info("requesting data from the database")
	if err != nil {
		return nil, err
	}
	s.logger.Info("request completed successfully")
	defer rows.Close()

	var bks api.BooksSet

	for rows.Next() {
		var bk api.Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Year)
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

func (s *GRPCServer) GetAuthors(ctx context.Context, book *api.Title) (*api.Authors, error) {
	var authors api.Authors

	rows, err := s.DB.Query("SELECT author FROM books, authors WHERE title = ? AND books.isbn = authors.isbn", book.Title)
	s.logger.Info("requesting data from the database")
	if err != nil {
		return nil, err
	}
	s.logger.Info("request completed successfully")
	defer rows.Close()

	for rows.Next() {
		var author string
		err := rows.Scan(&author)

		if err != nil {
			return nil, err
		}

		authors.Author = append(authors.Author, author)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &authors, err
}
