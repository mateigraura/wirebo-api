package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mateigraura/wirebo-api/utils"
	"time"
)

const expireTime = 604800

type JwtClaims struct {
	Id string
	jwt.StandardClaims
}

func GenerateJwt(id string) (string, error) {
	envVariables := utils.GetEnvFile()

	claims := JwtClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Unix() + expireTime,
			Issuer:    envVariables[utils.JWTIssuer],
		},
	}
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return payload.SignedString([]byte(envVariables[utils.JWTSecret]))
}

func ValidateJwt(signedToken string) (bool, error) {
	claims, err := parseToken(signedToken)
	if err != nil {
		return false, err
	}

	return isExpired(claims)
}

func GetClaims(signedToken string) (*JwtClaims, error) {
	return parseToken(signedToken)
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
		return nil, ErrJwtParse
	}

	return claims, nil
}

func isExpired(claims *JwtClaims) (bool, error) {
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return true, ErrJwtExpired
	}

	return true, nil
}
