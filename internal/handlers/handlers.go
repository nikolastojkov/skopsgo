package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/web/templates"
)

func LoadHandlers(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		comp := templates.Splash()
		templates.Layout(comp).Render(c, c.Writer)
	})

	loadCounterHandler(r)
}
