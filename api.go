package main

import (
	"github.com/gin-gonic/gin"
)

func loginAdminAPI(c *gin.Context) {
	email := c.Query("email")
	if email != "admin" {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not admin user",
		})
		return
	}
	password := c.Query("password")
	var user *User
	user = GetUserObj(email, password)
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
