package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
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
	router.POST("/api/addSite", addSiteAPI)
	router.POST("/api/loginAdmin", loginAdminAPI)
	router.POST("/api/regenerateSiteToken", regenerateSiteTokenAPI)
	router.GET("/api/getAllSites", getAllSitesAPI)
	router.POST("/api/userLoginByEmail", userLoginByEmailAPI)
	router.POST("/api/userRegisterByEmail", userRegisterByEmailAPI)

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
