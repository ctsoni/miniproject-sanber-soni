package repository

import (
	"database/sql"
	"log"
	"miniproject-sanber-soni/entity"
)

type CategoryRepository interface {
	GetCategory() ([]entity.Category, error)
	InsertCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(category entity.Category) (entity.Category, error)
	DeleteCategory(category entity.Category) error
	GetBookByCategoryId(category entity.Category) ([]entity.Book, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategory() ([]entity.Category, error) {
	var results []entity.Category
	sqlStatement := "SELECT * FROM category"
	data, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	for data.Next() {
		var category entity.Category

		err := data.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, category)
	}

	return results, nil
}

func (r *categoryRepository) InsertCategory(category entity.Category) (entity.Category, error) {
	sqlStatement := `
	INSERT INTO category (name, created_at, updated_at) 
	VALUES ($1, $2, $3)
	RETURNING *`
	err := r.db.QueryRow(
		sqlStatement,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(category entity.Category) (entity.Category, error) {
	sqlStatement := `
	UPDATE category
	SET name = $1, updated_at = $2 
	WHERE id = $3
	RETURNING *`
	err := r.db.QueryRow(
		sqlStatement,
		category.Name,
		category.UpdatedAt,
		category.ID).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(category entity.Category) error {
	sqlStatement := "DELETE FROM category WHERE id = $1"
	err := r.db.QueryRow(sqlStatement, category.ID)

	return err.Err()
}

func (r *categoryRepository) GetBookByCategoryId(category entity.Category) ([]entity.Book, error) {
	var books []entity.Book
	sqlStatement := `
	SELECT * FROM book
	WHERE category_id = $1`

	data, err := r.db.Query(sqlStatement, category.ID)
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
		books = append(books, book)
	}

	return books, nil
}
