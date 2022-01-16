package receipt

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterHandlers(rg *gin.RouterGroup, service Service) {
	receipts := rg.Group("/receipts")

	receipts.POST("", func(c *gin.Context) {
		var requestReceipt CreateReceiptRequest
		if err := c.ShouldBindJSON(&requestReceipt); err != nil {
			log.Panicln(" Bind json error. ")
		}
		log.Println(requestReceipt)
		service.Create(CreateReceiptRequest{
			Name:     requestReceipt.Name,
			Kind:     requestReceipt.Kind,
			Category: requestReceipt.Category,
			Contents: requestReceipt.Contents,
			Amount:   requestReceipt.Amount,
			Assets:   requestReceipt.Assets,
		})
		c.JSON(http.StatusCreated, nil)
	})

	receipts.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, service.Get(id))
	})

	receipts.PUT("", func(c *gin.Context) {
		var requestReceipt CreateReceiptRequest
		if err := c.ShouldBindJSON(&requestReceipt); err != nil {
			log.Panicln(" Bind json error. ")
		}
		log.Println(requestReceipt)
		service.Update(CreateReceiptRequest{
			Name:     requestReceipt.Name,
			Kind:     requestReceipt.Kind,
			Category: requestReceipt.Category,
			Contents: requestReceipt.Contents,
			Amount:   requestReceipt.Amount,
			Assets:   requestReceipt.Assets,
		})
		c.JSON(http.StatusCreated, nil)
	})
}
