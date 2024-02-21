package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/nikolastojkov/skopsgo/internal/config"
  "os"
)

func main() {
  config.LoadEnvConfig()

  port := os.Getenv("SERVER_PORT")

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.Run(fmt.Sprintf(":%s", port))
}
