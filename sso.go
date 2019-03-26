package main

import (
	"github.com/globalsign/mgo/bson"
)

type Site struct {
	SiteToken string
}

type User struct {
	Token    string
	Site     []string
	Password string
	Email    string
	Data     interface{}
}

// User declaration

func (u *User) save() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Insert(u)
}

func (u *User) update() {
	conn := mongoSession.Copy()
	defer conn.Close()
	_ = conn.DB("").C("user").Update(
		bson.M{"email": u.Email}, u)
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
func GetToken(email string, password string) *User {
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

func (u *User) Logout() {
	u.reassignToken()
}

func (u *User) AddSite(newSite string) {
	u.Site = append(u.Site, newSite)
	u.update()
}
