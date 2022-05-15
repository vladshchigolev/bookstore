package models

import (
	"database/sql"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

type Book struct {
	Isbn   int64
	Title  string
	Author string
	Year   uint16
}

// AllBooks returns a slice of all books in the books table.
func AllBooks() ([]Book, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book

	for rows.Next() {
		var bk Book

		err := rows.Scan(&bk.Title, &bk.Author, &bk.Year, &bk.Isbn)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
