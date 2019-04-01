package main

import (
	"github.com/globalsign/mgo/bson"
)

type Site struct {
	Domain           string
	SitePublicToken  string
	SitePrivateToken string
}

type SiteWithoutPrivateToken struct {
	Domain          string
	SitePublicToken string
}

type User struct {
	Token      string
	Password   string
	Identifier string
	Data       interface{}
	SiteTokens []string
}

type UserLog struct {
	Identifier string
	UserIP     string
	Action     int
	ActionTime string
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

func (ul *UserLog) save() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("log").Insert(ul)
}

func (u *User) update() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Update(
		bson.M{"email": u.Identifier}, u)
}

func (s *Site) update() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Update(
		bson.M{"domain": s.Domain}, s)
}

func (u *User) reassignToken() {
	var currentTime = getCurrentTime()
	u.Token = encryptString(u.Token + currentTime)
	u.update()
}

func RegisterByEmail(email string, password string, ip string) *User {
	var encryptedPassword = encryptString(password)
	var currentTime = getCurrentTime()
	var resultUser = &User{
		Password:   encryptedPassword,
		Identifier: email,
		Token:      encryptString(encryptedPassword + currentTime),
	}
	var log = &UserLog{
		Identifier: email,
		UserIP:     ip,
		Action:     0,
		ActionTime: currentTime,
	}
	resultUser.save()
	log.save()
	return resultUser
}

// aka Login
func GetUserObj(identifier string, password string, ip string) *User {
	var resultUser = &User{}
	var encryptedPassword = encryptString(password)
	var currentTime = getCurrentTime()
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"identifier": identifier,
			"password": encryptedPassword}).One(&resultUser)
	var log = &UserLog{
		Identifier: identifier,
		UserIP:     ip,
		Action:     1,
		ActionTime: currentTime,
	}
	log.save()
	return resultUser
}

func (u *User) isUserAdmin() bool {
	if u.Identifier == "admin" {
		return true
	}
	return false
}

func GetUserWithToken(token string) *User {
	var resultUser = &User{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"token": token}).One(&resultUser)
	return resultUser
}

func isTokenAdmin(token string) bool {
	var user *User
	user = GetUserWithToken(token)
	if user.Identifier == "admin" {
		return true
	}
	return false
}

// TODO: Testing
func GetUserWithSiteToken(siteToken string) *User {
	var resultUser = &User{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Find(
		bson.M{"$in": bson.M{"sitetoken": siteToken}}).One(&resultUser)
	return resultUser
}

func (u *User) Logout() {
	u.SiteTokens = nil
	u.reassignToken()
}

func (u *User) UserAddSiteToken(token string) {
	u.SiteTokens = append(u.SiteTokens, token)
	u.update()
}

func isSiteAdded(domain string) bool {
	conn := mongoSession.Copy()
	defer conn.Close()
	var site = &Site{}
	_ = conn.DB("").C("site").Find(
		bson.M{"domain": domain}).One(&site)
	return site.Domain != ""
}

func AddSite(domain string, key string) *Site {
	conn := mongoSession.Copy()
	defer conn.Close()
	var currentTime = getCurrentTime()
	var resultSite = &Site{
		Domain:           domain,
		SitePublicToken:  encryptString(domain + currentTime),
		SitePrivateToken: encryptString(domain + currentTime + key),
	}
	resultSite.save()
	return resultSite
}

func (s *Site) RegenerateSiteToken() {
	var currentTime = getCurrentTime()
	s.SitePublicToken = encryptString(s.SitePublicToken + currentTime)
	s.SitePrivateToken = encryptString(s.SitePrivateToken + currentTime)
	s.update()
}

func getAllSites() *[]SiteWithoutPrivateToken {
	var sites []SiteWithoutPrivateToken
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{}).All(&sites)
	return &sites
}

func getSiteByDomain(domain string) *Site {
	var resultSite = &Site{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{"domain": domain}).One(&resultSite)
	return resultSite
}

func getSiteByPublicToken(token string) *Site {
	var resultSite = &Site{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{"sitepublictoken": token}).One(&resultSite)
	return resultSite
}

func getSiteByPrivateToken(token string) *Site {
	var resultSite = &Site{}
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("site").Find(
		bson.M{"siteprivatetoken": token}).One(&resultSite)
	return resultSite
}

func LoginUserForSiteByEmail(token string, email string,
	password string, ip string) (string, string, string) {
	var site *Site
	var user *User
	var newSiteToken string
	site = getSiteByPublicToken(token)
	if "" != site.Domain {
		user = GetUserObj(email, password, ip)
		if "" != user.Token {
			newSiteToken = encryptString(user.Token + site.SitePrivateToken)
			user.UserAddSiteToken(newSiteToken)
			return user.Token, newSiteToken, site.Domain
		}
	}
	return "", "", ""
}

func isIdentifierUsed(identifier string) int {
	conn := mongoSession.Copy()
	defer conn.Close()
	count, _ := conn.DB("").C("user").Find(
		bson.M{"identifier": identifier}).Count()
	return count
}

func RegisterUserForSiteByEmail(token string, email string,
	password string, ip string) (string, string, string) {
	var site *Site
	var user *User
	var newSiteToken string
	site = getSiteByPublicToken(token)
	if "" != site.Domain {
		user = RegisterByEmail(email, password, ip)
		if "" != user.Token {
			newSiteToken = encryptString(user.Token + site.SitePrivateToken)
			user.UserAddSiteToken(newSiteToken)
			return user.Token, newSiteToken, site.Domain
		}
	}
	return "", "", ""
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
