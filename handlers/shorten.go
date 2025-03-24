package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"url-shortener/database"
	"url-shortener/utils"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
    OriginalURL string `json:"original_url"`
}

type ResponseBody struct {
    ShortCode string `json:"short_code"`
    ShortURL  string `json:"short_url"`
}

func ShortenURL(ctx *gin.Context){
	var reqBody RequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid JSON Request"})
        return
    }

	// validate url format
	if match, _ := regexp.MatchString(`^(https?:\/\/)?([\w\-]+\.)+[\w]{2,}(\/\S*)?$`, reqBody.OriginalURL); !match{
		ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid URL format"})
		return
	}
	
	// check in db for uniqueness
	var existingCode string
	err := database.DB.QueryRow("SELECT short_code FROM urls WHERE original_url=$1", reqBody.OriginalURL).Scan(&existingCode)
	if err == nil {
        // URL already exists, return the existing short code
        response := ResponseBody{
            ShortCode: existingCode,
            ShortURL:  fmt.Sprintf("http://localhost:8080/%v", existingCode),
        }
		ctx.JSON(http.StatusOK, response)
        return
    }

	// gen a short code
	// code := utils.GenerateShortCode(reqBody.OriginalURL)
	// shortenedURL := fmt.Sprintf("http://localhost:8080/%v", code)
	var code string
	for attempts := 0; attempts < 10; attempts++{ // putting in a loop to check its presence, and gen a diff code
												  // limit to 10 to prevent endless check
		code = utils.GenerateShortCode()
		var exists string 
		err := database.DB.QueryRow("SELECT short_code from urls WHERE short_code=$1", code).Scan(&exists)
		if err == sql.ErrNoRows {
            break // Found a unique code
        }
	}

	// add to db
	_, err = database.DB.Exec("INSERT INTO urls (original_url, short_code) VALUES($1, $2)", reqBody.OriginalURL, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "Database error"})
        return
	}
	
	// return the code and url
	response := ResponseBody{
		ShortCode: code,
		ShortURL: fmt.Sprintf("http://localhost:8080/%v", code),	
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}