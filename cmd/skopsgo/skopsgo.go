package main

import (
  "net/http"

  "log"
  "fmt"
  "os"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  port := os.Getenv("SERVER_PORT")

  if port == "" {
    port = "8080"
  }

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.Run(fmt.Sprintf(":%s", port))
}
