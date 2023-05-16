package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const jwtSecret = "rabndom secret key"

func GenerateToken(email string, role string) (string, error) {
	tokenLifespanHrs := 1

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespanHrs)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func IsTokenValid(tokenString string, role string) bool {
	token, err := parseToken(tokenString)
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims["role"] == role {
		return true
	}

	return false
}
func ValidateRequestToken(c *gin.Context, role string) bool {
	tokenString := extractToken(c)
	return IsTokenValid(tokenString, role)
}

func ExtractTokenEmail(c *gin.Context) (string, error) {
	tokenString := extractToken(c)
	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%s", claims["email"]), nil
	}
	return "", nil
}

func extractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token not signed properly")
		}
		return []byte(jwtSecret), nil
	})
	return token, err
}
