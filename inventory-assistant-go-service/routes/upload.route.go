package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadRegister(route *gin.Engine) {
	route.POST("/upload", func(c *gin.Context) {

		file, error := c.FormFile("file")
		if error != nil {
			fmt.Println("", error)
		}

		wd, err := os.Getwd()
		c.SaveUploadedFile(file, filepath.Join(wd, file.Filename))

		if err != nil {
			fmt.Println("", err)
		}

		c.String(http.StatusOK, "success")
	})
}
