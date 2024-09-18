package security

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/magiconair/properties"
)

var jwtSecret string

const PropertyFilePath = "/properties/api.properties"

func getJwtSecret() string {
	if jwtSecret == "" {
		wd, err := os.Getwd()
		fmt.Println(wd)
		if err != nil {
			panic(err)
		}
		props := properties.MustLoadFile(wd + PropertyFilePath, properties.UTF8)
		jwtSecret = props.GetString("authentication.api.jwtSecret", "test")
	}
	return jwtSecret
}



// public method CreateJWTToken creates a new JWT token with the given usernam.
func CreateJWTToken(username string) (string, error) {
	return createJWTToken(jwt.MapClaims{"username": username})
}

// private method createJWTToken creates a new JWT token with the given claims and secret key.
func createJWTToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := getJwtSecret()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateJWTToken validates the given JWT token using the secret key.
func ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(getJwtSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}