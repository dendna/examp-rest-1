package bookRepository

import (
	"database/sql"
	"log"

	"github.com/dendna/examp-rest-1/models"
)

// BookRepository ...
type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks ...
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select id, title, author, year from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	return books
}

// GetBook ...
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select id, title, author, year from books where id = $1", id)
	_ = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	//logFatal(err)

	return book
}

// AddBook ...
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) returning id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)
	logFatal(err)

	return book.ID
}

// UpdateBook ...
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	res, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 returning id;",
		&book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)

	rowsUpdated, err := res.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

// RemoveBook ...
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	res, err := db.Exec("delete from books where id = $1", id)
	logFatal(err)

	rowsDeleted, err := res.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
