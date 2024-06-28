package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Serve static files
	r.Static("/static", "./web/static")

	r.LoadHTMLGlob("web/*.html")

	r.GET("/", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Hello, World!",
		// })

		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
