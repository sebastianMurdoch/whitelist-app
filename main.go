package main

import (
	"fmt"
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
		var data map[string]interface{}
		if err := c.Bind(&data); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(data)
		c.JSON(http.StatusOK, data)
	})

	router.Run(":" + port)
}
