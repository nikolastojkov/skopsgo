package handlers

import (
	"github.com/gin-gonic/gin"
  "fmt"
) 

var counterValue int = 1

func loadCounterHandler(router *gin.Engine) {
  router.PUT("/counterIncrement", func(context *gin.Context) {
		counterValue++
		context.String(context.Writer.Status(), fmt.Sprintf("%d", counterValue))
	})

	router.PUT("/counterDecrement", func(context *gin.Context) {
		if counterValue != 1 {
			counterValue--
		}

		context.String(context.Writer.Status(), fmt.Sprintf("%d", counterValue))
	})
}
