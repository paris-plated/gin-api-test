package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/paris-plated/gin-api-test/requests"
)

func Heroku(c *gin.Context) {
	fmt.Println("PARAMS: %v", c.Request.URL.Query())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var webhook requests.Webhook
	if err := c.BindJSON(&webhook); err != nil {
		// heroku doesn't care about this response so we should just
		// log / notify ourselves that there was an error parsing

		return
	}

	// This auth shit should be middlewear
	var auth_token string

	if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {
		auth_token = strings.Split(values[0], " ")[1]
	} else {
		auth_token = ""
	}

	if auth_token != os.Getenv("HEROKU_SECRET") {
		// // heroku won't do anything if the request is unauthorized so we
		// // should also just log / notify ourselves that something was fishy

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong_token",
			"token":   auth_token,
		})
		return
	}

	// Send webhook to circle

	// Heroku perfers this status since everything is ignored
	// c.JSON(http.StatusNoContent, gin.H{
	c.JSON(http.StatusOK, gin.H{
		"action": webhook.Action,
	})
}
