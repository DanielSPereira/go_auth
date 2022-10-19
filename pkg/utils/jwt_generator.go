package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func GenerateTokens(userId string) (*Tokens, error) {
	// create accessToken and return if fail
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return nil, err
	}

	// create refreshToken and return if fail
	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// return tokens if success
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GenerateAccessToken(userId string) (string, error) {
	// create claims for the token
	claims := jwt.MapClaims{}

	// get expiration_in_milliseconds
	expirationInMilliseconds, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRATION_IN_MILISECONDS"))

	// define claims for the token
	claims["id"] = userId
	claims["expires"] = time.Now().Add(time.Millisecond * time.Duration(expirationInMilliseconds))

	// get JWT secret
	jwtSecret := os.Getenv("secret")

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	accessToken, err := token.SignedString([]byte(jwtSecret))

	return accessToken, err
}

func GenerateRefreshToken() (string, error) {
	hash := sha256.New()

	refresh := os.Getenv("REFRESH_TOKEN_SECRET") + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	expirationInMilliseconds, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_IN_MILISECONDS"))

	expirationTime := fmt.Sprint(time.Now().Add(time.Millisecond * time.Duration(expirationInMilliseconds)).Unix())

	refreshToken := hex.EncodeToString(hash.Sum(nil)) + "." + expirationTime

	return refreshToken, nil
}
