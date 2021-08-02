package main

import (
	"fmt"

	"github.com/felipefinhane/go-crud-gin-gorm/config"
	"github.com/felipefinhane/go-crud-gin-gorm/models"
	"github.com/felipefinhane/go-crud-gin-gorm/routes"
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
