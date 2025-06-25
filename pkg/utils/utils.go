package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	newtoken, err := token.SignedString([]byte("secret"))
	return "Authorization " + newtoken, err
}
func CheckPassword(pwd string, hashpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pwd))
	return err == nil
}
func CheckJwt(tokenString string) (string, error) {
	if len(tokenString) > 14 && tokenString[:14] == "Authorization " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	// validate the essential claims
	if !token.Valid {
		return "", errors.New("token验证不通过")
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if name, ok := claims["username"].(string); ok {
				return name, nil
			} else {
				return "", errors.New("token不包含username")
			}
		}
	}
	return "", errors.New("token解析出现错误")
}
