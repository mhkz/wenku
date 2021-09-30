package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wenku/model/dbModel"
	"wenku/model/modelInterface"
)

type RegistStuct struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

func Regist(c *gin.Context) {
	var R RegistStuct
	_ = c.BindJSON(&R)
	U := dbModel.NewUser(dbModel.User{UserName: R.UserName, PassWord: R.PassWord})
	var curd modelInterface.CURD
	curd = U
	err, user := curd.Create()
	fmt.Println(err, user)
}
