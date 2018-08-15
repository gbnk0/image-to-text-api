package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
)

func getText(image []byte) string {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(image)
	client.Languages = []string{"eng", "fra"}
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
		version = strings.TrimSuffix(version, "\n")
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
	host := flag.String("h", "0.0.0.0", "Set port for listening")
	port := flag.Int("p", 8000, "Set port for listening")
	flag.Parse()

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

	bindString := *host + ":" + strconv.Itoa(*port)
	r.Run(bindString)
}
