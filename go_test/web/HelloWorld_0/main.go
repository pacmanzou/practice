package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	defer db.Close()
	router := gin.Default()
	router.POST("/api/auth/register", func(c *gin.Context) {
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
			name = RandomStringrange(10)
		}

		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newUser)
		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"message": "register successfully",
		})
	})
	router.Run()
}
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "goWeb"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("falied to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

const letterBytes = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func RandomStringrange(n int) string {
	res := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range res {
		res[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(res)
}
