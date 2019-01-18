package main

import (
  "github.com/gin-gonic/gin"
  "github.com/paris-plated/gin-api-test/handlers"
)

func main() {
  router := gin.Default()

  router.GET("/", handlers.Heroku)
  router.POST("/", handlers.Heroku)

  router.Run()
}
