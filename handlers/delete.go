package handlers

import (
	"net/http"
	"url-shortener/database"
	"github.com/gin-gonic/gin"
)

func DeleteShortURL(ctx *gin.Context){
	shortCode := ctx.Param("shortcode")

	res, err := database.DB.Exec("DELETE FROM urls WHERE short_code=$1", shortCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0{
		ctx.JSON(http.StatusNotFound, gin.H{"error" : "Short URL not found"})
        return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message" : "URL deleted successfully"})
}