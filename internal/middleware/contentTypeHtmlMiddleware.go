package middleware

import (
  "github.com/gin-gonic/gin"
)

func ContentTypeHtmlMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
    context.Writer.Header().Set("Content-Type", "text/html charset=utf-8")
	}
}
