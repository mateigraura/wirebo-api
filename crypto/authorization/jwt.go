package authorization

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mateigraura/wirebo-api/core/utils"
	"github.com/mateigraura/wirebo-api/crypto"
)

type JwtClaims struct {
	Id string
	jwt.StandardClaims
}

var isExpired = func(claims JwtClaims) bool {
	return claims.ExpiresAt < time.Now().UTC().Unix()
}

func GenerateJwt(id string) (string, error) {
	envVariables := utils.GetEnvFile()
	minutesToExpiration, _ := strconv.Atoi(envVariables[utils.JWTExpiry])
	claims := JwtClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * time.Duration(minutesToExpiration)).Unix(),
			Issuer:    envVariables[utils.JWTIssuer],
		},
	}
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return payload.SignedString([]byte(envVariables[utils.JWTSecret]))
}

func ValidateJwt(signedToken string) (JwtClaims, error) {
	claims, err := parseToken(signedToken)
	if err != nil {
		return JwtClaims{}, err
	}

	if isExpired(*claims) {
		return JwtClaims{}, crypto.ErrJwtExpired
	}

	return *claims, nil
}

func GetClaims(signedToken string, verify bool) (JwtClaims, error) {
	var claims *JwtClaims
	var err error
	if !verify {
		claims, err = parseTokenUnverified(signedToken)
	} else {
		claims, err = parseToken(signedToken)
	}

	if err != nil {
		return JwtClaims{}, err
	}
	return *claims, nil
}

func parseToken(signedToken string) (*JwtClaims, error) {
	envVariables := utils.GetEnvFile()

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(envVariables[utils.JWTSecret]), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, crypto.ErrJwtParse
	}

	return claims, nil
}

func parseTokenUnverified(signedToken string) (*JwtClaims, error) {
	envVariables := utils.GetEnvFile()

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(envVariables[utils.JWTSecret]), nil
		},
	)

	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&(jwt.ValidationErrorExpired) != 0 && token != nil {
				claims, ok := token.Claims.(*JwtClaims)
				if !ok {
					return nil, crypto.ErrJwtParse
				}

				return claims, nil
			}
		}

		return nil, err
	}

	return nil, crypto.ErrJwtParse
}
