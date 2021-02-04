package controller

import (
	"net/http"
	"pacmanzou/HelloWorld/common"
	"pacmanzou/HelloWorld/model"
	"pacmanzou/HelloWorld/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	// 获取参数
	db := common.GetDB()
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// 密码,手机号格式验证
	if !regularExpressionValidator(c, password, phone) {
		return
	}

	// 判断手机号是否存在
	if isPhoneExist(db, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "该用户已经存在",
		})
		return
	}

	// 如果密码为空，默认位随机的10位字符串
	if name == "" {
		name = util.RandomString(10)
	}

	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "加密错误",
		})
		return
	}

	// 初始化
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}

	// 向数据表插入数据
	db.Create(&newUser)

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"message": "register successfully",
	})
}

func Login(c *gin.Context) {

	// 获取参数
	db := common.GetDB()
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	// 密码,手机号格式验证
	if !regularExpressionValidator(c, password, phone) {
		return
	}

	// 判断手机号是否存在
	if !isPhoneExist(db, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "该用户不存在",
		})
		return
	}

	// 判断密码是否正确
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "密码输入错误",
		})
		return
	}

	// 发放token
	token := "11"

	// 返回结果
	c.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "login successfully",
	})
}

func regularExpressionValidator(c *gin.Context, password string, phone string) bool {
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return false
	}

	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号要为11位",
		})
		return false
	}
	return true
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
