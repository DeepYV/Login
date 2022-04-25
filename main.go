package main

import (
	"loginapi/login/controller"
	"loginapi/login/password"
	"loginapi/login/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Assign() user.Key {

	a, b := password.GenerateKey()
	K := user.Key{}

	K.Pkey = a
	K.Pubkey = b
	return K

}

func main() {

	router := gin.Default()
	key := Assign()
	user.KeyValue(key)
	router.POST("/SS", controller.SignUp)
	router.POST("/User", controller.LoginApi)

	router.Run("localhost:8080")

}
