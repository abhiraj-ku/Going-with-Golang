package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env load
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env file")
	}

	// router config for gin
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 8 << 20 //8 MiB

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from gin",
		})
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")

		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "failed to upload file",
			})
			return
		}
		// save the file
		err = c.SaveUploadedFile(file, "assets/uploads"+file.Filename)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "failed to upload file",
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/assets/uploads" + file.Filename,
		})

	})

	router.Run()
}
