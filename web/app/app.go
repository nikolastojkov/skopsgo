package app

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
  "os"
)

func LoadApp() {
  port := os.Getenv("SERVER_PORT")
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.Run(fmt.Sprintf(":%s", port))
}
