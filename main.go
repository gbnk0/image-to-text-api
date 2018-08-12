package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
)

func getText(image []byte) string {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(image)
	text, _ := client.Text()
	return text
}

func getVersion() string {
	versionFile := "version.txt"
	version := "dev"
	if _, err := os.Stat(versionFile); err == nil {

		b, err := ioutil.ReadFile(versionFile)
		if err != nil {
			fmt.Print(err)
		}
		version = string(b)
	}
	return version
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r := gin.Default()

	r.POST("/text", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		check(err)
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			check(err)
			return
		}

		text := getText(buf.Bytes())
		c.JSON(200, gin.H{
			"status": "success",
			"text":   text,
		})

	})

	r.GET("/version", func(c *gin.Context) {
		version := getVersion()
		c.JSON(200, gin.H{
			"status":  "success",
			"version": version,
		})

	})
	r.Run(":8000")
}
