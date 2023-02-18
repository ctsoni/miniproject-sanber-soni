package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// accounts admin-password dan editor-secret
// post put delete wajib

func BasicAuth(ctx *gin.Context) {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal(err)
	}

	user, pwd, hasAuth := ctx.Request.BasicAuth()
	if hasAuth && user == os.Getenv("USER_ONE") && pwd == os.Getenv("PWD_ONE") {
		log.Println("User Authenticated")
	} else if hasAuth && user == os.Getenv("USER_TWO") && pwd == os.Getenv("PWD_TWO") {
		log.Println("User Authenticated")
	} else {
		log.Println("User not Authenticated")
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "You're not authorized to do this method",
		})
		ctx.Abort()
		return
	}
}
