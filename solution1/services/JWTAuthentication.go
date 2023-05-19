package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type CustomClaim struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

type JWTService interface {
	GenerateToken(username string, isAdmin bool, exp time.Duration) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "AuthService",
	}
}

func (j *jwtService) GenerateToken(username string, isAdmin bool, exp time.Duration) string {
	//TODO implement me
	claim := &jwt.MapClaims{
		"iss":      j.issuer,
		"exp":      time.Now().Add(time.Minute * exp).Unix(),
		"username": username,
		"is_admin": isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, _ := token.SignedString([]byte(j.secretKey))

	return signedToken
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	//TODO implement me
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func getSecretKey() string {
	//TODO implement me
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "secret"
	}
	return secret
}
