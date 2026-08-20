package utils

import "time"
import "github.com/golang-jwt/jwt/v5"

var jwtSecret = []byte("MYSTORE_SECRET_KEY_CHANGE_IN_PRODUCTION")

// GenerateToken creates a JWT signed token containing user ID
func GenerateToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}