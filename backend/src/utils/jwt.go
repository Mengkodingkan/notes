package util

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func GenerateToken(Data map[string]interface{}, SecretKeyEnv string) (string, error) {
	// expiredAt := time.Now().Add(time.Duration(time.Minute) * (100000 * 1537)).Unix()

	jwtSecretKey := Get(SecretKeyEnv)
	claims := jwt.MapClaims{}
	claims["ExpiredAt"] = 10
	claims["Authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())

		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(c *gin.Context, SecretKeyEnv string) (*jwt.Token, error) {
	tokenHeader := c.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	secretKey := Get(SecretKeyEnv)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())

		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := Get(SecrePublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
