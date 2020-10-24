package service

import (
	"os"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

func Encode(userId string) string {
	signKey := os.Getenv("PRIVATE_KEY")
	// fmt.Println(signKey)
	// verifyKey := os.Getenv("PUBLIC_KEY")
	// fmt.Println(verifyKey)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UserID"] 	= userId
	claims["exp"] 		= time.Now().Add(time.Hour * 72).Unix()
    claims["iat"] 		= time.Now().Unix()

	tokenString, err := token.SignedString([]byte(signKey))
	fmt.Println("token: ", tokenString, "\nerror: ", err)

	return tokenString
}

func Decode(token string) string {
	return ""
}