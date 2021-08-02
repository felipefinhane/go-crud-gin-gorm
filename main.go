package main

import (
	"fmt"

	"github.com/felipefinhane/go-crud-gin-gorm/config"
	"github.com/felipefinhane/go-crud-gin-gorm/models"
	"github.com/felipefinhane/go-crud-gin-gorm/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := config.DBURL(config.BuildDBConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status: ", err)
	}
	config.DB = db

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	//Running
	r.Run()
}
