package main

import (
	"github.com/globalsign/mgo"
)

const url = "127.0.0.1:27017/sso"

var mongoSession, err = mgo.Dial(url)

const salt = "5ad5f3a339b335625fa56507b080bec468c34fcc30ad09894c44aa6dd4b13435"
const ADMIN_EMAIL = "admin@admin.com"
