package app

import (
	"fmt"
	"github.com/nikolastojkov/skopsgo/web/templates"
	"os"

	"github.com/gin-gonic/gin"
)

func LoadApp() {
	port := os.Getenv("SERVER_PORT")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		comp := templates.Hello("World")
		comp.Render(c, c.Writer)
	})

	r.Run(fmt.Sprintf(":%s", port))
}
