package models

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func InitializeJwtSecret() (err error) {

	jwtSecret := RandomString(256)

	_, err = SetConfig("jwt_secret", jwtSecret)
	if err != nil {
		return err
	}

	return nil
}

func InitializePassword() (err error) {

	password := RandomString(32)

	_, err = SetConfig("password", password)
	if err != nil {
		return err
	}

	return nil
}

func CheckPassword(password string) (correct bool, err error) {

	passwordConfig, err := GetConfig("password")
	if err != nil {
		return false, err
	}

	serverPassword := passwordConfig.Value

	return password == serverPassword, nil
}

func IssueJwt(username string) (tokenString string, err error) {

	jwtSecretConfig, err := GetConfig("jwt_secret")
	if err != nil {
		return "", err
	}

	// Define the JWT claims
	claims := &Claims{
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err = token.SignedString([]byte(jwtSecretConfig.Value))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ValidateJwt(tokenString string) (bool, error) {
	// Retrieve the JWT secret from the configuration
	jwtSecretConfig, err := GetConfig("jwt_secret")
	if err != nil {
		return false, err
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the secret key
		return []byte(jwtSecretConfig.Value), nil
	})

	// Check if parsing and validation were successful
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
