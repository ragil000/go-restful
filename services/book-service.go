package services

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ragil000/go-restful.git/dto"
	"github.com/ragil000/go-restful.git/entities"
	"github.com/ragil000/go-restful.git/repositories"
)

// BookService is a...
type BookService interface {
	Insert(b dto.BookCreateDTO) entities.Book
	Update(b dto.BookUpdateDTO) entities.Book
	Delete(b entities.Book)
	All() []entities.Book
	FindByID(bookID uint64) entities.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repositories.BookRepository
}

// NewBookService is a...
func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entities.Book {
	book := entities.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDTO) entities.Book {
	book := entities.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entities.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entities.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entities.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	fmt.Println("data: ")
	fmt.Println(b)
	fmt.Println("UserID: ")
	fmt.Println(userID)
	fmt.Println("ID: ")
	fmt.Println(id)
	return userID == id
}
