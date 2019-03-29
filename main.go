package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"APIHost":   "http://127.0.0.1:8080/",
		"CrawlHost": "http://127.0.0.1:8848/",
	})
}

func register(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func admin(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{})
}

func startServer() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//load html templates
	router.LoadHTMLGlob("templates/*")

	//api

	//web pages
	router.GET("/", login)
	router.GET("/register", register)
	router.GET("/admin", admin)

	router.Static("/assets", "./assets")
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Run(":8080")
}

func main() {
	startServer()
}
