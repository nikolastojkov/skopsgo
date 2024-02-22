package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/internal/handlers"
	"os"
)

func LoadApp() {
	port := os.Getenv("SERVER_PORT")
	r := gin.Default()
	r.Static("/static", "./web/static")

	// Load middleware
	handlers.LoadHandlers(r)

	r.Run(fmt.Sprintf(":%s", port))
}
