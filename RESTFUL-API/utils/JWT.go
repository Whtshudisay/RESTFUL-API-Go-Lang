package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "SUPANIGGA"

func GenerateToken(email string, userid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	Parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing Method !")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not Parse token")
	}

	tokenIsValid := Parsedtoken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid Token !")
	}

	claims, ok := Parsedtoken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid Token Claims")
	}

	//email:= claims["email"].(string)
	userid := int64(claims["userid"].(float64))
	return userid, nil
}
