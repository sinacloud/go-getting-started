package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5050"
	}

	r := gin.Default()
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/files", "./files")
	r.LoadHTMLGlob("templates/*.tmpl.html")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ping.tmpl.html", nil)
	})

	r.Run(":" + port)
}
