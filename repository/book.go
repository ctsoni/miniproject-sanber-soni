package repository

import (
	"database/sql"
	"log"
	"miniproject-sanber-soni/entity"
)

type BookRepository interface {
	GetBook() ([]entity.Book, error)
	InsertBook(book entity.Book) (entity.Book, error)
	UpdateBook(book entity.Book) (entity.Book, error)
	DeleteBook(book entity.Book) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) GetBook() ([]entity.Book, error) {
	var result []entity.Book
	sqlStatement := "SELECT * FROM book"
	data, err := b.db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	for data.Next() {
		var book entity.Book
		err := data.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.CategoryID)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, book)
	}

	return result, nil
}

func (b *bookRepository) InsertBook(book entity.Book) (entity.Book, error) {
	sqlStatement := `
	INSERT INTO book (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING *`
	err := b.db.QueryRow(
		sqlStatement,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CreatedAt,
		book.UpdatedAt,
		book.CategoryID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.CategoryID)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *bookRepository) UpdateBook(book entity.Book) (entity.Book, error) {
	sqlStatement := `
	UPDATE book 
	SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, updated_at = $9 
	WHERE id = $10
	RETURNING *`

	err := b.db.QueryRow(
		sqlStatement,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		book.UpdatedAt,
		book.ID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.CategoryID)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *bookRepository) DeleteBook(book entity.Book) error {
	sqlStatement := "DELETE FROM book WHERE id = $1"
	err := b.db.QueryRow(sqlStatement, book.ID)

	if err != nil {
		return err.Err()
	}

	return nil
}
