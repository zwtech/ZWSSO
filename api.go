package main

import (
	"github.com/gin-gonic/gin"
)

func loginAdminAPI(c *gin.Context) {
	un := c.Query("email")
	if un != "admin" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not admin user",
		})
		return
	}
	password := c.Query("password")
	ip := c.ClientIP()
	var user *User
	user = GetUserObj(un, password, ip)
	if user.isUserAdmin() {
		c.JSON(200, gin.H{
			"success": 1,
			"token":   user.Token,
			"info":    "",
		})
	} else {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not admin user or password incorrect",
		})
	}
}

func addSiteAPI(c *gin.Context) {
	var token string
	token = c.Query("token")
	if token == "" {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "No token provided",
		})
		return
	}
	if !isTokenAdmin(token) {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "User is not admin",
		})
		return
	}
	var domain string
	domain = c.Query("domain")
	if domain == "" {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "No domain provided",
		})
		return
	}
	if isSiteAdded(domain) {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "Site is added",
		})
		return
	}
	var site *Site
	site = AddSite(domain, token)
	c.JSON(200, gin.H{
		"success":          1,
		"sitePublicToken":  site.SitePublicToken,
		"sitePrivateToken": site.SitePrivateToken,
		"info":             "",
	})
}

func regenerateSiteTokenAPI(c *gin.Context) {
	var token string
	token = c.Query("token")
	if token == "" {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "No token provided",
		})
		return
	}
	if !isTokenAdmin(token) {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "User is not admin",
		})
		return
	}
	var domain string
	domain = c.Query("domain")
	if domain == "" {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "No domain provided",
		})
		return
	}
	if !isSiteAdded(domain) {
		c.JSON(200, gin.H{
			"success":          0,
			"sitePublicToken":  "",
			"sitePrivateToken": "",
			"info":             "No site found",
		})
		return
	}
	var site *Site
	site = getSiteByDomain(domain)
	site.RegenerateSiteToken()
	c.JSON(200, gin.H{
		"success":          1,
		"sitePublicToken":  site.SitePublicToken,
		"sitePrivateToken": site.SitePrivateToken,
		"info":             "",
	})
}

func getAllSitesAPI(c *gin.Context) {
	var token string
	token = c.Query("token")
	if token == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"sites":   nil,
			"info":    "No token provided",
		})
		return
	}
	if !isTokenAdmin(token) {
		c.JSON(200, gin.H{
			"success": 0,
			"sites":   nil,
			"info":    "User is not admin",
		})
		return
	}
	var sites *[]SiteWithoutPrivateToken
	sites = getAllSites()

	c.JSON(200, gin.H{
		"success": 1,
		"sites":   sites,
		"info":    "",
	})
}

func userLoginByEmailAPI(c *gin.Context) {
	var publicToken string
	publicToken = c.Query("publicToken")
	if publicToken == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not a public token",
		})
		return
	}
	email := c.Query("email")
	if email == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "No email",
		})
		return
	}
	password := c.Query("password")
	if password == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "No password",
		})
		return
	}
	ip := c.ClientIP()
	var userSiteToken string
	userSiteToken = LoginUserForSiteByEmail(publicToken, email, password, ip)
	if userSiteToken == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not a valid email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": 1,
		"token":   userSiteToken,
		"info":    "",
	})
}

func userRegisterByEmailAPI(c *gin.Context) {
	var publicToken string
	publicToken = c.Query("publicToken")
	if publicToken == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not a public token",
		})
		return
	}
	email := c.Query("email")
	if email == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "No email",
		})
		return
	}
	password := c.Query("password")
	if password == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "No password",
		})
		return
	}
	ip := c.ClientIP()
	var userSiteToken string
	userSiteToken = RegisterUserForSiteByEmail(publicToken, email, password, ip)
	if userSiteToken == "" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not a valid email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": 1,
		"token":   userSiteToken,
		"info":    "",
	})
}
