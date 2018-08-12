package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"

	"net/http"
	"os"
	// "path/filepath"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		text := getText(file.Filename)
		c.JSON(200, gin.H{
			"status": "success",
			"text":   text,
		})

		os.Remove(file.Filename)
	})
	r.Run(":8000")
}

func getText(imagePath string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(imagePath)
	text, _ := client.Text()
	return text
}
