package main

import (
	"fmt"

	"git@github.com/felipefinhane/go-crud-gin-gorm.git/config"
	"git@github.com/felipefinhane/go-crud-gin-gorm.git/models"
	"git@github.com/felipefinhane/go-crud-gin-gorm.git/routes"
)

func main() {
	config.DB, err = gorm.Opem("mysql", config.DBURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	//Running
	r.Runt()
}
