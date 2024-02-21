package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nikolastojkov/skopsgo/web/templates"
)

var count int

func LoadHandlers(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		comp := templates.Splash()
		templates.Layout(comp).Render(c, c.Writer)
	})

	// TODO: Transfer htmx handlers elsewhere, a.k.a. refactor handlers package
	r.POST("/counterIncrement", func(c *gin.Context) {
		count++
		c.String(c.Writer.Status(), fmt.Sprintf("%d", count))
	})

	r.POST("/counterDecrement", func(c *gin.Context) {
		if count != 0 {
			count--
		}
		c.String(c.Writer.Status(), fmt.Sprintf("%d", count))
	})

}
