package user

import (
	"crypto/rsa"
)

type LoginUserRequest struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccesstokenExpires  int64
	RefreshTokenExpires int64
	UserId              string
}

type Key struct {
	Pkey   *rsa.PrivateKey
	Pubkey rsa.PublicKey
}

type LoginApi struct {
	Name     string `gorm:"unique" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

var k Key

func KeyValue(key Key) {
	k = key

}
func AssignStruct() Key {

	return k

}
