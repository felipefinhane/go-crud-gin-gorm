package config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

//DBConfig represents db configuration
type DBConfig struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DBName string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:   "172.18.0.3",
		Port:   3306,
		User:   "root",
		Pass:   "root",
		DBName: "gorm",
	}

	return &dbConfig

}

func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
