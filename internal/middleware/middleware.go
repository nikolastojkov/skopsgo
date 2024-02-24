package middleware

import (
  "github.com/gin-gonic/gin"
)

func LoadMiddleware(router *gin.Engine) {
	router.Use(ContentTypeHtmlMiddleware())
}
