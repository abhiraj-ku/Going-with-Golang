// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// .env load
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error loading env file")
// 	}

// 	// router config for gin
// 	router := gin.Default()
// 	router.Static("/assets", "./assets")
// 	router.LoadHTMLGlob("templates/*")
// 	router.MaxMultipartMemory = 8 << 20 // 8 MiB

// 	// setup s3 uploader
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		log.Fatalf("unable to load SDK config, %v", err)
// 		return
// 	}

// 	// setup aws client setup using cfg
// 	client := s3.NewFromConfig(cfg)

// 	// uploader using the client
// 	uploader := manager.NewUploader(client)

// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello from gin",
// 		})
// 	})
// 	router.GET("/home", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{})
// 	})

// 	// save to the local disk
// 	// router.POST("/upload", func(c *gin.Context) {
// 	// 	file, err := c.FormFile("image")
// 	// 	if err != nil {
// 	// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 	// 			"error": "failed to upload file",
// 	// 		})
// 	// 		return
// 	// 	}
// 	// 	// save the file
// 	// 	err = c.SaveUploadedFile(file, "assets/uploads"+file.Filename)
// 	// 	if err != nil {
// 	// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 	// 			"error": "failed to upload file",
// 	// 		})
// 	// 		return
// 	// 	}
// 	// })

// 	// Push the file to AWS S3 bucket API
// 	router.POST("/upload", func(c *gin.Context) {
// 		file, err := c.FormFile("image")
// 		if err != nil {
// 			c.HTML(http.StatusOK, "index.html", gin.H{
// 				"error": "failed to upload the file",
// 			})
// 			return
// 		}

// 		// open the file to send through the Body
// 		f, openErr := file.Open()
// 		if openErr != nil {
// 			c.HTML(http.StatusOK, "index.html", gin.H{
// 				"error": "Error while opening",
// 			})
// 			return
// 		}
// 		defer f.Close() // Ensure the file is closed after handling

// 		// save this file to AWS S3 bucket
// 		result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 			Bucket: aws.String("go-corona"),
// 			Key:    aws.String(file.Filename),
// 			Body:   f,
// 			// ACL:    "public-read",
// 		})

// 		if uploadErr != nil {
// 			// c.HTML(http.StatusOK, "index.html", gin.H{
// 			// 	"error": "some error on AWS",
// 			// })
// 			// return
// 			c.JSON(500, gin.H{
// 				"error": "error on Aws upload",
// 			})
// 		}

// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 			"image": result.Location,
// 		})
// 	})

// 	router.Run()
// }

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	// Initialize the Gin router
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Setup S3 uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create S3 client from configuration
	client := s3.NewFromConfig(cfg)

	// Create an uploader with the S3 client
	uploader := manager.NewUploader(client)

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
				"error": "failed to upload the file",
			})
			return
		}

		// Open the file to send through the Body
		f, openErr := file.Open()
		if openErr != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "failed to upload the file",
			})
			return
		}
		defer f.Close() // Ensure the file is closed after handling

		// Save this file to AWS S3 bucket
		result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("go-corona"),
			Key:    aws.String(file.Filename),
			Body:   f,
			ACL:    types.ObjectCannedACLPublicRead,
		})

		if uploadErr != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "some error on AWS",
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": result.Location,
		})
	})

	router.Run()
}
