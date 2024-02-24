package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/internal/handlers"
	"github.com/nikolastojkov/skopsgo/internal/middleware"
	"os"
)

func LoadApp() {
	port := os.Getenv("SERVER_PORT")
	router := gin.Default()
	router.Static("/static", "./web/static")

	middleware.LoadMiddleware(router)
	handlers.LoadHandlers(router)

	router.Run(fmt.Sprintf(":%s", port))
}
