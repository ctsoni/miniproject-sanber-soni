package controllers

import (
	"github.com/gin-gonic/gin"
	"miniproject-sanber-soni/entity"
	"miniproject-sanber-soni/helper"
	"miniproject-sanber-soni/service"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (b *bookHandler) GetBook(ctx *gin.Context) {
	var result gin.H
	books, err := b.bookService.GetBook()
	if err != nil {
		result = gin.H{
			"message": err,
		}
		return
	} else {
		result = gin.H{
			"result": books,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (b *bookHandler) InsertBook(ctx *gin.Context) {
	var inputBook entity.InputBook

	err := ctx.ShouldBindJSON(&inputBook)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot proccess the request",
			"error":   helper.FormatError(err),
		})
		return
	}

	newBook, err := b.bookService.InsertBook(inputBook)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "There is no such category_id (foreign key) in category table",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mesasge": "insert success",
		"result":  newBook,
	})
}

func (b *bookHandler) UpdateBook(ctx *gin.Context) {
	var inputBook entity.InputBook
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot parse id into integer",
		})
		return
	}

	err = ctx.ShouldBindJSON(&inputBook)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot proccess the request",
			"error":   helper.FormatError(err),
		})
		return
	}

	newBook, err := b.bookService.UpdateBook(inputBook, id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "error mapping input to book struct",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mesasge": "update success",
		"result":  newBook,
	})
}

func (b *bookHandler) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "cannot parse id into integer",
		})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "error deleting row",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
