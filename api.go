package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func loginAdminAPI(c *gin.Context) {
	un := c.PostForm("email")
	fmt.Println(un)
	if un != ADMIN_EMAIL {
		c.JSON(200, gin.H{
			"success": 0,
			"token":   "",
			"info":    "Not admin user",
		})
		return
	}
	password := c.PostForm("password")
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
			"info":    "Password incorrect",
		})
	}
}

func addSiteAPI(c *gin.Context) {
	var token string
	token = c.PostForm("token")
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
	domain = c.PostForm("domain")
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
	token = c.PostForm("token")
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
	domain = c.PostForm("domain")
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
	publicToken = c.PostForm("publicToken")
	if publicToken == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "Not a public token",
		})
		return
	}
	email := c.PostForm("email")
	if email == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "No email",
		})
		return
	}
	password := c.PostForm("password")
	if password == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "No password",
		})
		return
	}
	ip := c.ClientIP()
	var userSiteToken string
	var domain string
	var token string
	token, userSiteToken, domain = LoginUserForSiteByEmail(publicToken, email, password, ip)
	if userSiteToken == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "Not a valid email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"success":   1,
		"siteToken": userSiteToken,
		"token":     token,
		"domain":    domain,
		"info":      "",
	})
}

func userRegisterByEmailAPI(c *gin.Context) {
	var publicToken string
	publicToken = c.PostForm("publicToken")
	if publicToken == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "Not a public token",
		})
		return
	}
	email := c.PostForm("email")
	if email == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "No email",
		})
		return
	}
	password := c.PostForm("password")
	if password == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "No password",
		})
		return
	}
	ip := c.ClientIP()
	var userSiteToken string
	var domain string
	var token string
	token, userSiteToken, domain = RegisterUserForSiteByEmail(publicToken, email, password, ip)
	if userSiteToken == "" {
		c.JSON(200, gin.H{
			"success":   0,
			"siteToken": "",
			"token":     "",
			"domain":    "",
			"info":      "Not a valid email or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"success":   1,
		"siteToken": userSiteToken,
		"token":     token,
		"domain":    domain,
		"info":      "",
	})
}
