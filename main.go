package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"miniproject-sanber-soni/controllers"
	"miniproject-sanber-soni/database"
	"miniproject-sanber-soni/middleware"
	"miniproject-sanber-soni/repository"
	"miniproject-sanber-soni/service"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"miniproject")

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection to database is failed")
		panic(err)
	}

	fmt.Println("Successfully make connection to database")

	database.DbMigrate(db)
	defer db.Close()

	r := gin.Default()

	// bangun datar
	bangunDatar := r.Group("/bangun-datar")
	bangunDatar.GET("/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	bangunDatar.GET("/persegi", controllers.Persegi)
	bangunDatar.GET("/persegi-panjang", controllers.PersegiPanjang)
	bangunDatar.GET("/lingkaran", controllers.Lingkaran)

	// category and book
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := controllers.InitCategoryHandler(categoryService)
	categories := r.Group("/categories")
	categories.GET("", categoryHandler.GetCategory)
	categories.POST("", middleware.BasicAuth, categoryHandler.InsertCategory)
	categories.PUT("/:id", middleware.BasicAuth, categoryHandler.UpdateCategory)
	categories.DELETE("/:id", middleware.BasicAuth, categoryHandler.DeleteCategory)
	categories.GET("/:id/books", categoryHandler.GetBookByCategoryId)

	// book
	bookRepository := repository.NewBookRepo(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := controllers.NewBookHandler(bookService)
	book := r.Group("/books")
	book.GET("", bookHandler.GetBook)
	book.POST("", middleware.BasicAuth, bookHandler.InsertBook)
	book.PUT("/:id", middleware.BasicAuth, bookHandler.UpdateBook)
	book.DELETE("/:id", bookHandler.DeleteBook)

	r.Run()
}
