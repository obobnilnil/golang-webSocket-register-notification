package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateTokenI(email string) (string, error) {
	var jwtSecret = []byte("your_jwt_secret")
	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = email
	claims["requires_action"] = "change_password" // fixed the typo and added this field to the token
	claims["exp"] = expirationTime.Unix()         // Token expiration time

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
