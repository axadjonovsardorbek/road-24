package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "mrbek"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJWTToken(userID string, username string, role string, is_admin string) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["username"] = username
	claims["role"] = role
	claims["is_admin"] = is_admin
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(180 * time.Minute).Unix() // Token expires in 3 hours
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token : ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["username"] = username
	claims["role"] = role
	claims["is_admin"] = is_admin
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(24 * time.Hour).Unix() // Refresh token expires in 24 hours
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func GenerateJWTTokenForDriver(role string, car_number string, technical_passport string, is_admin string, id string) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["role"] = role
	claims["car_number"] = car_number
	claims["technical_passport"] = technical_passport
	claims["is_admin"] = is_admin
	claims["user_id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(180 * time.Minute).Unix() // Token expires in 3 hours
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token : ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["role"] = role
	claims["car_number"] = car_number
	claims["technical_passport"] = technical_passport
	claims["is_admin"] = is_admin
	claims["id"] = id
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(24 * time.Hour).Unix() // Refresh token expires in 24 hours
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	fmt.Print(token.Claims)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
