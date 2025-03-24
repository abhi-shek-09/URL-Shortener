package handlers

import (
	"database/sql"
	"net/http"
	"url-shortener/database"

	"github.com/gin-gonic/gin"
)

func RetrieveURL(ctx *gin.Context){
	shortCode := ctx.Param("shortcode")
	var originalURL string
	err := database.DB.QueryRow("SELECT original_url FROM urls WHERE short_code=$1", shortCode).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, gin.H{"error" : "Short URL not found"})
            return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "Database error"})
        return
	}

	_, err = database.DB.Exec("UPDATE urls SET clicks=clicks + 1 WHERE short_code=$1", shortCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to update clicks"})
        return
	}

	ctx.Redirect(http.StatusSeeOther, originalURL)
}