package controller

import (
	"net/http"
	"pacmanzou/HelloWorld/common"
	"pacmanzou/HelloWorld/model"
	"pacmanzou/HelloWorld/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号要为11位",
		})
		return
	}
	if isPhoneExist(db, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "该用户已经存在",
		})
		return
	}
	if name == "" {
		name = util.RandomString(10)
	}

	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"message": "register successfully",
	})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
