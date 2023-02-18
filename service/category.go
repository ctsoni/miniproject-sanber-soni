package service

import (
	"miniproject-sanber-soni/entity"
	"miniproject-sanber-soni/repository"
	"time"
)

type CategoryService interface {
	GetCategory() ([]entity.Category, error)
	InsertCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(category entity.Category, id int) (entity.Category, error)
	DeleteCategory(id int) error
	GetBookByCategoryId(id int) ([]entity.Book, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) GetCategory() ([]entity.Category, error) {
	cat, err := c.categoryRepo.GetCategory()
	if err != nil {
		return cat, err
	}

	return cat, nil
}

func (c *categoryService) InsertCategory(category entity.Category) (entity.Category, error) {
	var cat entity.Category

	cat.Name = category.Name
	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()

	newCat, err := c.categoryRepo.InsertCategory(cat)
	if err != nil {
		return newCat, err
	}

	return newCat, nil
}

func (c *categoryService) UpdateCategory(category entity.Category, id int) (entity.Category, error) {
	var cat entity.Category

	cat.ID = id
	cat.Name = category.Name
	cat.UpdatedAt = time.Now()

	newCat, err := c.categoryRepo.UpdateCategory(cat)
	if err != nil {
		return newCat, err
	}

	return newCat, err
}

func (c *categoryService) DeleteCategory(id int) error {
	var cat entity.Category

	cat.ID = id
	err := c.categoryRepo.DeleteCategory(cat)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryService) GetBookByCategoryId(id int) ([]entity.Book, error) {
	var cat entity.Category

	cat.ID = id
	books, err := c.categoryRepo.GetBookByCategoryId(cat)
	if err != nil {
		return books, err
	}

	return books, err
}
