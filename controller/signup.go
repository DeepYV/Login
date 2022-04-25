package controller

import (
	"encoding/base64"
	"fmt"
	connect "loginapi/login/Connect"
	"loginapi/login/password"
	dd "loginapi/login/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	db1 := connect.GetDatabase()

	var user dd.LoginApi
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//encrypt
	PP := dd.AssignStruct()

	pass := password.RSA_Encrypt([]byte(user.Password), PP.Pkey)

	name := user.Name
	email := user.Email
	country := user.Country
	phone := user.Phone
	password := base64.StdEncoding.EncodeToString(pass)
	Value := &dd.LoginApi{Name: name, Email: email, Country: country, Phone: phone, Password: password}

	result := db1.Table("loginapi").Create(Value)
	//	insertStmt := `insert into "LoginApi"("name", "email", "country","phone","password") values($1, $2, $3, $4,$5)`
	//	_, e := db1.Exec(insertStmt, name, email, country, phone, password)
	if result.Error != nil {

		panic(result.Error)
	}
	fmt.Println("accept")
	c.JSON(http.StatusAccepted, gin.H{"ok": "accepted"})

}
