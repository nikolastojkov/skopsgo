package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/internal/handlers"
	"os"
)

func LoadApp() {
	port := os.Getenv("SERVER_PORT")
	router := gin.Default()
	router.Static("/static", "./web/static")

	// Load middleware
	handlers.LoadHandlers(router)

	router.Run(fmt.Sprintf(":%s", port))
}
