package utility

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/deckarep/golang-set"
	"github.com/golang-jwt/jwt"
)

var (
	jwtSecretKey = GetToken()
	validToken   = mapset.NewSet()
)

func GetToken() string {
	hash := sha256.New()
	temp := make([]byte, 64)
	rand.Read(temp)
	hash.Write(temp)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetJWT(user string) (string, error) {
	claims := jwt.StandardClaims{Audience: user, ExpiresAt: time.Now().Unix() + 600, Id: user, IssuedAt: time.Now().Unix(), Issuer: "CNLib", NotBefore: time.Now().Unix(), Subject: "Token"}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	validToken.Add(token)
	return token, nil
}

func ValidateJWT(token string) bool {
	ok := validToken.Contains(token)
	if !ok {
		return false
	}
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(jwtSecretKey), nil })
	return err == nil && jwtToken.Valid
}

func ExpiredJWT(token string) {
	// Suppose the token is valid
	validToken.Remove(token)
}
