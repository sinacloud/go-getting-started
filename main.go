package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5050"
	}

	db, err := sql.Open("mysql", os.ExpandEnv("$ACCESSKEY:$SECRETKEY@tcp($MYSQL_HOST:$MYSQL_PORT)/app_$APPNAME"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("templates/*.tmpl.html")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ping.tmpl.html", nil)
	})
	r.GET("/mysql", func(c *gin.Context) {
		var result int

		if err := db.QueryRow("SELECT 1 + 1").Scan(&result); err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("SELECT 1 + 1 = %v", result))
	})

	r.Run(":" + port)
}
