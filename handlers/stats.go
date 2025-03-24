package handlers

import (
	"net/http"
	"url-shortener/database"
	"github.com/gin-gonic/gin"
)

type Stats struct {
	ShortCode string `json:"short_code"`
	Clicks    int    `json:"access_count"`
}

func GenerateStats(ctx *gin.Context) {
	shortCode := ctx.Param("shortcode")

	var count int
	err := database.DB.QueryRow("SELECT clicks FROM urls WHERE short_code=$1", shortCode).Scan(&count)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error" : "Short URL not found"})
		return
	}

	stats := Stats{
		ShortCode: shortCode,
		Clicks:    count,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, stats)
}
