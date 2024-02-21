package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func LoadApp() {
	port := os.Getenv("SERVER_PORT")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.Run(fmt.Sprintf(":%s", port))
}
