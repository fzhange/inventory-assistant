package main

import (
	"fmt"
	"inventory-assistant/routes"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:Pa123457@tcp(localhost:3306)/hello-mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		fmt.Println("", err)
	} else {
		fmt.Println("", db)
	}
}
