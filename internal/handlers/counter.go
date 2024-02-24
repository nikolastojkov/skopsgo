package handlers

import (
	"github.com/gin-gonic/gin"
  "fmt"
) 

var counterValue int = 1

func loadCounterHandler(router *gin.Engine) {
  counterRouter := router.Group("/counter")

  counterRouter.PUT("/increment", func(context *gin.Context) {
		counterValue++
		context.String(context.Writer.Status(), fmt.Sprintf("%d", counterValue))
	})

	counterRouter.PUT("/decrement", func(context *gin.Context) {
		if counterValue != 1 {
			counterValue--
		}

		context.String(context.Writer.Status(), fmt.Sprintf("%d", counterValue))
	})
}
