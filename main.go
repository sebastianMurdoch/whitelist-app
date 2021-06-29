package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	
	router.POST("/data", func(c *gin.Context) {
		var data []byte
		if err := c.Bind(&data); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		log.Println(string(data))
		c.String(http.StatusOK, string(data))
	})

	router.Run(":" + port)
}
