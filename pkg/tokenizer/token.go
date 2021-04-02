package tokenizer

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/scys12/modul-go/internal/config"
)

var cfg config.Config

func init() {
	cfg, _ = config.InitConfig()
}

func CreateToken(role string, userid int) (string, error) {

	atClaims := jwt.MapClaims{
		"role":       role,
		"user_id":    userid,
		"exp":        time.Now().Add(time.Minute * time.Duration(cfg.JWTExpires)).Unix(),
		"authorized": true,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func ExtractToken(tokenStr string) string {
	var accessToken string
	if tokenStr[:6] != "Bearer" {
		accessToken = ""
	} else {
		tokenArr := strings.Split(tokenStr, " ")
		if len(tokenArr) == 2 {
			accessToken = tokenArr[1]
		}
	}
	return accessToken
}

func DecodeToken(token string) (map[string]interface{}, error) {
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signed methods")
		}
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok && !tokenParsed.Valid {
		return nil, err
	}
	tokenDetails := make(map[string]interface{})
	tokenDetails["user_id"] = claims["user_id"]
	tokenDetails["role"] = claims["role"]
	return tokenDetails, nil
}
