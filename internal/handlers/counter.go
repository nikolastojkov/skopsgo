package handlers

import (
	"github.com/gin-gonic/gin"
  "fmt"
) 

var count int = 1

func loadCounterHandler(r *gin.Engine) {
  r.POST("/counterIncrement", func(c *gin.Context) {
		count++
		c.String(c.Writer.Status(), fmt.Sprintf("%d", count))
	})

	r.POST("/counterDecrement", func(c *gin.Context) {
		if count != 1 {
			count--
		}

		c.String(c.Writer.Status(), fmt.Sprintf("%d", count))
	})
}
