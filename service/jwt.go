package service

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SignKey 	string
	VerifyKey 	string
}

func NewJWT(signKey string, verifyKey string) *JWT {
	jwt := JWT{signKey, verifyKey}
	return &jwt
}

func (j *JWT) Encode(payload string) (error, string) {
	token 			   := jwt.New(jwt.SigningMethodHS256)
	claims 			   := token.Claims.(jwt.MapClaims)
	claims["UserID"] 	= payload
	claims["exp"] 		= time.Now().Add(time.Hour * 72).Unix()
    claims["iat"] 		= time.Now().Unix()
    tokenString, err   := token.SignedString([]byte(j.SignKey))
    
    return err, tokenString
}