package main

import (
	"github.com/globalsign/mgo"
	"github.com/go-redis/redis"
)

const url = "127.0.0.1:27017/sso"

var mongoSession, err = mgo.Dial(url)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       3,
})

const salt = "5ad5f3a339b335625fa56507b080bec468c34fcc30ad09894c44aa6dd4b13435"
