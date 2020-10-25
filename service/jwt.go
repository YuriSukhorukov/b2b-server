package service

import (
	"time"
	// "fmt"
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
	claims["payload"] 	= payload
	claims["exp"] 		= time.Now().Add(time.Hour * 72).Unix()
    claims["iat"] 		= time.Now().Unix()
    tokenString, err   := token.SignedString([]byte(j.SignKey))

    j.Decode(tokenString)
    
    return err, tokenString
}

func (j *JWT) Decode(t string) (error, string) {
	tokenString := t  
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	    return []byte(j.SignKey), nil
	})
	// ... error handling

	// do something with decoded claims
	// for key, val := range claims {
	//     fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	// fmt.Println(claims["payload"])
	payload := claims["payload"].(string)

	return err, payload
}