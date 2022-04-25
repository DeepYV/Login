package controller

import (
	"encoding/base64"
	connect "loginapi/login/Connect"
	"loginapi/login/password"
	"loginapi/login/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginApi(c *gin.Context) {

	var UserCred user.LoginUserRequest

	if err := c.ShouldBindJSON(&UserCred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection := connect.GetDatabase()
	var result user.LoginApi

	connection.Raw("SELECT name,password FROM loginapi where name=?", UserCred.Username).Scan(&result)

		PP := user.AssignStruct()
		DS, _ := base64.StdEncoding.DecodeString(result.Password)

		Verify := password.Verify([]byte(UserCred.Password), []byte(DS), PP.Pubkey)
		if Verify == true {

			c.JSON(http.StatusOK, gin.H{"message": "login"})
			token, err := CreateToken(UserCred.Username)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}
			c.JSON(http.StatusOK, token)
			return
		} else {

			c.JSON(http.StatusBadRequest, gin.H{"message": "authentication failed"})
		}

	}

