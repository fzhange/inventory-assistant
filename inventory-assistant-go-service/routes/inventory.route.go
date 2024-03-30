package routes

import (
	"inventory-assistant/tools"

	"github.com/gin-gonic/gin"
)

func inventoryRegister(r *gin.Engine) {
	r.GET("/match-order", func(ctx *gin.Context) {
		cookedOrderData := tools.CookOrderData()

		ctx.JSON(200, gin.H{
			"cookedOrderData": cookedOrderData,
		})
	})
}
