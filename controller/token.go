package controller

import (
	"loginapi/login/user"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (user.TokenDetails, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = userId
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return user.TokenDetails{}, err

	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return user.TokenDetails{}, err
	}
	tk := user.TokenDetails{

		AccessToken:         t,
		RefreshToken:        rt,
		AccesstokenExpires:  time.Now().Add(time.Minute * 15).Unix(),
		RefreshTokenExpires: time.Now().Add(time.Hour * 24).Unix(),
		UserId:              userId,
	}
	return tk, nil

}
