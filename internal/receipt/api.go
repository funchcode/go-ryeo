package receipt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlers(rg *gin.RouterGroup) {
	receipts := rg.Group("/receipts")

	receipts.POST("", func(c *gin.Context) {
		NewService().Create(CreateReceiptRequest{})
		c.JSON(http.StatusCreated, nil)
	})

	receipts.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello. chae hwan")
	})
}
