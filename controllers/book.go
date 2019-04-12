package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dendna/examp-rest-1/models"
	bookRepository "github.com/dendna/examp-rest-1/repository/book"
	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks ...
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		repo := bookRepository.BookRepository{}
		books = repo.GetBooks(db, book, books)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	}
}

// GetBook ...
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)

	}
}

// AddBook ...
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		bookID = bookRepo.AddBook(db, book)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookID)

	}
}

// UpdateBook ...
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

// RemoveBook ...
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := bookRepo.RemoveBook(db, id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
