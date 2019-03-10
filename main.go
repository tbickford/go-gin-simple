package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.Printf("Application is running")
	r.GET("/ping", func(c *gin.Context) {

		log.Println("ping called")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post-test", func(c *gin.Context) {
		payload := map[string]string{}
		c.Bind(&payload)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  payload["message"],
		})
	})

	r.GET("/test-params/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Query("name")
		c.String(http.StatusOK, fmt.Sprintf("id [%v], name [%v]", id, name))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
