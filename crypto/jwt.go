package crypto

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mateigraura/wirebo-api/utils"
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
		return JwtClaims{}, ErrJwtExpired
	}

	return *claims, nil
}

func GetClaims(signedToken string) (JwtClaims, error) {
	claims, err := parseToken(signedToken)
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
		return nil, ErrJwtParse
	}

	return claims, nil
}
