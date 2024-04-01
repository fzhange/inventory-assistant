package main

import (
	"fmt"
	"inventory-assistant/model"
	"inventory-assistant/routes"
	"inventory-assistant/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	routes.RegisterRoutes(engine)

	test()

	// listen 8080 port
	engine.Run()
}

func test() {
	db, _ := tools.SetupDB()

	var users []model.User
	db.Model(&model.User{}).Preload("Post").Find(&users)
	fmt.Println("", users)
}
