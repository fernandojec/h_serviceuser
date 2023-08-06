package auths

import "github.com/golang-jwt/jwt/v4"

var APPLICATION_NAME = "MyApplication"
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("4p1Management")
var ACCESS_SECRET = "4p1Management"
var REFRESH_SECRET = "4p1Management"
