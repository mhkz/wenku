package main

import (
	"wenku/config"
	"wenku/init/initRouter"
	"wenku/init/mysql"
	"wenku/init/registTable"
	"net/http"
	"time"
)

func main() {

	registTable.RegistTable(mysql.InitMysql(config.Dbconfig.Admin))
	defer mysql.DEFAULTDB.Close()
	Router := initRouter.InitRouter()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}