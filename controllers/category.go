package controllers

import (
	"github.com/gin-gonic/gin"
	"miniproject-sanber-soni/entity"
	"miniproject-sanber-soni/service"
	"net/http"
	"strconv"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func InitCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (c *categoryHandler) GetCategory(ctx *gin.Context) {
	var result gin.H
	cat, err := c.categoryService.GetCategory()
	if err != nil {
		result = gin.H{
			"message": err,
		}
		return
	} else {
		result = gin.H{
			"result": cat,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *categoryHandler) InsertCategory(ctx *gin.Context) {
	var category entity.Category

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot proccess entity",
		})
		return
	}

	newCat, err := c.categoryService.InsertCategory(category)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot make new category",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "insert category success",
		"result":  newCat,
	})
}

func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var cat entity.Category
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot parse id into string",
		})
		return
	}

	err = ctx.ShouldBindJSON(&cat)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot bind to json",
		})
		return
	}

	newCat, err := c.categoryService.UpdateCategory(cat, id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot update category",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"result":  newCat,
	})
}

func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot parse id into string",
		})
		return
	}

	err = c.categoryService.DeleteCategory(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot delete category",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
}

func (c *categoryHandler) GetBookByCategoryId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Cannot parse id into string",
		})
		return
	}

	var result gin.H
	books, err := c.categoryService.GetBookByCategoryId(id)
	if err != nil {
		result = gin.H{
			"message": "no books with that category_id or category_id not found in category table",
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	ctx.JSON(http.StatusOK, result)
}
