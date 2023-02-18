package service

import (
	"miniproject-sanber-soni/entity"
	"miniproject-sanber-soni/helper"
	"miniproject-sanber-soni/repository"
	"time"
)

type BookService interface {
	GetBook() ([]entity.Book, error)
	InsertBook(inputBook entity.InputBook) (entity.Book, error)
	UpdateBook(inputBook entity.InputBook, id int) (entity.Book, error)
	DeleteBook(id int) error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) *bookService {
	return &bookService{bookRepository}
}

func (b *bookService) GetBook() ([]entity.Book, error) {
	books, err := b.bookRepository.GetBook()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (b *bookService) InsertBook(inputBook entity.InputBook) (entity.Book, error) {
	var book entity.Book

	book.Title = inputBook.Title
	book.Description = inputBook.Description
	book.ImageURL = inputBook.ImageUrl
	book.ReleaseYear = inputBook.ReleaseYear
	book.Price = inputBook.Price
	book.TotalPage = inputBook.TotalPage
	book.CategoryID = inputBook.CategoryId

	book.Thickness = helper.FormatThickness(book.TotalPage)

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	newBook, err := b.bookRepository.InsertBook(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (b *bookService) UpdateBook(inputBook entity.InputBook, id int) (entity.Book, error) {
	var book entity.Book

	book.ID = id

	book.Title = inputBook.Title
	book.Description = inputBook.Description
	book.ImageURL = inputBook.ImageUrl
	book.ReleaseYear = inputBook.ReleaseYear
	book.Price = inputBook.Price
	book.TotalPage = inputBook.TotalPage
	book.CategoryID = inputBook.CategoryId

	book.Thickness = helper.FormatThickness(book.TotalPage)
	book.UpdatedAt = time.Now()

	newBook, err := b.bookRepository.UpdateBook(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (b *bookService) DeleteBook(id int) error {
	var book entity.Book
	book.ID = id

	err := b.bookRepository.DeleteBook(book)
	if err != nil {
		return err
	}

	return nil
}
