package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func learnRegister(r *gin.Engine) {
	v1 := r.Group("/learn")
	v1.GET("/login", loginEndpoint)
}

func loginEndpoint(c *gin.Context) {
	b := make(map[string]int)

	b["沙河娜扎"] = 100
	fmt.Println(b)

	c.JSON(200, gin.H{
		"message": "learn",
	})
}
