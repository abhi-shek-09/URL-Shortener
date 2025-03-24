package handlers

import (
	"net/http"
	"regexp"
	"url-shortener/database"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
    NewURL string `json:"new_url"`
}

func UpdateURL(ctx *gin.Context){
	shortCode := ctx.Param("shortcode")

	var reqBody UpdateRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid JSON Request"})
		return
	}

	if match, _ := regexp.MatchString(`^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(/.*)?$`, reqBody.NewURL); !match {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid URL format"})
        return
    }

	res, err := database.DB.Exec("UPDATE urls SET original_url = $1 WHERE short_code = $2", reqBody.NewURL, shortCode)
	// using $ helps prevent SQL injection
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "Database error"})
        return
	}

	// very important 
	rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error" : "Short URL not found"})
        return
    }

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, map[string]string {"message" : "Updated the url successfully"})
}