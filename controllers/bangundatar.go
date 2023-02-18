package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"strconv"
)

func SegitigaSamaSisi(ctx *gin.Context) {
	metodaHitung := ctx.Query("hitung")
	alasString := ctx.Query("alas")
	tinggiString := ctx.Query("tinggi")

	alas, err := strconv.Atoi(alasString)
	if err != nil {
		log.Fatal(err.Error())
	}

	tinggi, err := strconv.Atoi(tinggiString)
	if err != nil {
		log.Fatal(err.Error())
	}

	if metodaHitung == "luas" {
		luas := 0.5 * float64(alas*tinggi)
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan luas": luas,
		})
		return
	} else if metodaHitung == "keliling" {
		keliling := 3 * alas
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan keliling": keliling,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "parameter url tidak sesuai",
	})
}

func Persegi(ctx *gin.Context) {
	metodaHitung := ctx.Query("hitung")
	sisiString := ctx.Query("sisi")

	sisi, err := strconv.Atoi(sisiString)
	if err != nil {
		log.Fatal(err.Error())
	}

	if metodaHitung == "luas" {
		luas := sisi * sisi
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan luas": luas,
		})
		return
	} else if metodaHitung == "keliling" {
		keliling := 4 * sisi
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan keliling": keliling,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "parameter url tidak sesuai",
	})
}

func PersegiPanjang(ctx *gin.Context) {
	metodaHitung := ctx.Query("hitung")
	panjangString := ctx.Query("panjang")
	lebarString := ctx.Query("lebar")

	panjang, err := strconv.Atoi(panjangString)
	if err != nil {
		log.Fatal(err.Error())
	}

	lebar, err := strconv.Atoi(lebarString)
	if err != nil {
		log.Fatal(err.Error())
	}

	if metodaHitung == "luas" {
		luas := panjang * lebar
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan luas": luas,
		})
		return
	} else if metodaHitung == "keliling" {
		keliling := 2*panjang + 2*lebar
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan keliling": keliling,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "parameter url tidak sesuai",
	})
}

func Lingkaran(ctx *gin.Context) {
	metodaHitung := ctx.Query("hitung")
	jariJariString := ctx.Query("jariJari")

	jariJari, err := strconv.Atoi(jariJariString)
	if err != nil {
		log.Fatal(err.Error())
	}

	if metodaHitung == "luas" {
		luas := math.Pi * float64(jariJari*jariJari)
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan luas": luas,
		})
		return
	} else if metodaHitung == "keliling" {
		keliling := 2.0 * math.Pi * float64(jariJari)
		ctx.JSON(http.StatusOK, gin.H{
			"hasil perhitungan keliling": keliling,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "parameter url tidak sesuai",
	})
}
