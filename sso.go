package main

import (
	"github.com/globalsign/mgo/bson"
)

type Site struct {
	Domain           string
	SitePublicToken  string
	SitePrivateToken string
}

type User struct {
	Token      string
	Site       []string
	Password   string
	Email      string
	Data       interface{}
	SiteTokens []string
}

// User declaration

func (u *User) save() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Insert(u)
}

func (s *Site) save() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Insert(s)
}

func (u *User) update() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Update(
		bson.M{"email": u.Email}, u)
}

func (s *Site) update() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Update(
		bson.M{"email": s.SitePublicToken}, s)
}

func (u *User) reassignToken() {
	var currentTime = string(getCurrentTime())
	u.Token = encryptString(u.Token + currentTime)
	u.update()
}

func Register(email string, password string, data interface{}) *User {
	var encryptedPassword = encryptString(password)
	var currentTime = string(getCurrentTime())
	var resultUser = &User{
		Password: encryptedPassword,
		Email:    email,
		Data:     data,
		Token:    encryptString(encryptedPassword + currentTime),
	}
	resultUser.save()
	return resultUser
}

// aka Login
func GetUserObj(email string, password string) *User {
	var resultUser = &User{}
	var encryptedPassword = encryptString(password)
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"email": email,
			"password": encryptedPassword}).One(&resultUser)
	resultUser.reassignToken()
	return resultUser
}

func GetUserWithToken(token string) *User {
	var resultUser = &User{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"token": token}).One(&resultUser)
	return resultUser
}

// TODO: Testing
func GetUserWithSiteToken(siteToken string) *User {
	var resultUser = &User{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"$in": bson.M{"siteToken": siteToken}}).One(&resultUser)
	return resultUser
}

func (u *User) Logout() {
	u.reassignToken()
}

func (u *User) UserAddSite(newSite string) {
	u.Site = append(u.Site, newSite)
	u.update()
}

func (u *User) UserAddSiteToken(token string) {
	u.SiteTokens = append(u.SiteTokens, token)
	u.update()
}

func AddSite(domain string, key string) *Site {
	conn := mongoSession.Copy()
	defer conn.Close()
	var currentTime = string(getCurrentTime())
	var resultSite = &Site{
		Domain:           domain,
		SitePublicToken:  encryptString(domain + currentTime),
		SitePrivateToken: encryptString(domain + currentTime + key),
	}
	resultSite.save()
	return resultSite
}

func getSiteByPublicToken(token string) *Site {
	var resultSite = &Site{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{"sitePublicToken": token}).One(&resultSite)
	return resultSite
}

func getSiteByPrivateToken(token string) *Site {
	var resultSite = &Site{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{"sitePrivateToken": token}).One(&resultSite)
	return resultSite
}

func LoginUserForSite(token string, email string, password string) string {
	var site *Site
	var user *User
	var newSiteToken string
	site = getSiteByPublicToken(token)
	if "" != site.Domain {
		user = GetUserObj(email, password)
		if "" != user.Token {
			user.UserAddSite(site.SitePrivateToken)
			newSiteToken = encryptString(user.Token + site.SitePrivateToken)
			user.UserAddSiteToken(newSiteToken)
			return newSiteToken
		}
	}
	return ""
}

func getUserFromSiteToken(sitePrivateToken string, userSiteToken string) *User {
	var user *User
	var site *Site

	site = getSiteByPrivateToken(sitePrivateToken)
	if "" != site.Domain {
		user = GetUserWithSiteToken(userSiteToken)
		if "" != user.Token {
			if encryptString(user.Token+sitePrivateToken) == userSiteToken {
				return user
			}
		}
	}
	return nil
}
