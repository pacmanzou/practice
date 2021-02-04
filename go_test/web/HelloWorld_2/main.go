package main

import (
	"pacmanzou/HelloWorld/common"
	"pacmanzou/HelloWorld/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	router.Register(r)
	router.Login(r)
	r.Run()
}
