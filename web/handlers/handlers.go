package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/web/templates"
)

func LoadHandlers(r *gin.Engine) {
  r.GET("/", func(c *gin.Context) {
		comp := templates.RenderSplash()
		comp.Render(c, c.Writer)
	})

  r.GET("/getDone", func(c *gin.Context) {
		comp := templates.Done()
		comp.Render(c, c.Writer)
	})
}
