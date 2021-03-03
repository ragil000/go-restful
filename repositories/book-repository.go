package repositories

import (
	"fmt"

	"github.com/ragil000/go-restful.git/entities"
	"gorm.io/gorm"
)

// BookRepository is a...
type BookRepository interface {
	InsertBook(b entities.Book) entities.Book
	UpdateBook(b entities.Book) entities.Book
	DeleteBook(b entities.Book)
	AllBook() []entities.Book
	FindBookByID(bookID uint64) entities.Book
}

type bookConnection struct {
	connection *gorm.DB
}

// NewBookRepository is a...
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) InsertBook(b entities.Book) entities.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) DeleteBook(b entities.Book) {
	db.connection.Delete(&b)
}

func (db *bookConnection) UpdateBook(b entities.Book) entities.Book {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *bookConnection) FindBookByID(bookID uint64) entities.Book {
	fmt.Println("bookIDMM: ")
	fmt.Println(bookID)
	var book entities.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) AllBook() []entities.Book {
	var books []entities.Book
	db.connection.Preload("User").Find(&books)
	return books
}
