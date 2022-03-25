package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	db, err := gorm.Open("mysql", "root:21mysql@/gintest?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&HelloWorld{})

	// db.Create(&HelloWorld{
	// 	Name: "rui",
	// 	Age:  80,
	// 	Sex:  false,
	// })

	var hello []HelloWorld
	// db.Where("name = ?", "grant").First(&hello)
	// db.Where("age < ?", "25").Find(&hello)
	// db.Where("name = ?", "rui").First(&HelloWorld{}).Update("name", "newrui")

	// 软删除
	// db.Where("id = ?", 3).Delete(&HelloWorld{})

	// db.Where("id > ?", 3).Unscoped().Delete(&HelloWorld{})

	db.Find(&hello)
	fmt.Println(&hello)
}
