package main

import (
	"fmt"
	"log"
	"time"

	"github.com/daymenu/school/app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.SchoolInfo{})
	schoolInfo := model.SchoolInfo{}
	schoolInfo.Name = "吕梁学院"
	schoolInfo.CreatedAt = time.Now()
	schoolInfo.UpdatedAt = time.Now()

	db.Create(&schoolInfo)
	db.First(&schoolInfo)
	fmt.Println(schoolInfo)
	defer db.Close()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
