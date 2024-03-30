/*
fp: fmt.Println
frr: range
*/

package routes

import (
	"github.com/gin-gonic/gin"
)

func indexRegister(route *gin.Engine) {
	route.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "okppLLL;;;;;LLp",
		})
	})
}

func RegisterRoutes(route *gin.Engine) {
	indexRegister(route)
	learnRegister(route)
	uploadRegister(route)
	inventoryRegister(route)
}
