package utils

import (
	"authRestApis/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userEmail string) (string, string, error) {

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["user_email"] = userEmail
	accessTokenClaims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", "", fmt.Errorf("error signing access token: %w", err)
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["user_email"] = userEmail
	refreshTokenClaims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return "", "", fmt.Errorf("error signing refresh token: %w", err)
	}

	return signedAccessToken, signedRefreshToken, nil
}

func ValidateToken(tokenString, secretKey string) (string, error) {
	models.BlacklistedTokens.Lock()
	if exp, found := models.BlacklistedTokens.TokenMap[tokenString]; found && exp.Before(time.Now()) {
		models.BlacklistedTokens.Unlock()
		return "", errors.New("token has been revoked")
	}
	models.BlacklistedTokens.Unlock()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what we expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			if time.Now().After(expirationTime) {
				return "", fmt.Errorf("token has expired")
			}
		}
		if userID, ok := claims["user_email"].(string); ok {
			return userID, nil
		}
	}

	return "", fmt.Errorf("invalid token")
}
