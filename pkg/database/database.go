package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	//"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbConnection := viper.GetString("Database.DB_Connection")
	dbHost := viper.GetString("Database.DB_Host")
	dbPort := viper.GetInt("Database.DB_Port")
	dbName := viper.GetString("Database.DB_Name")
	dbUsername := viper.GetString("Database.DB_Username")
	dbPassword := viper.GetString("Database.DB_Password")
	dbTimezone := viper.GetString("Database.DB_TIMEZONE")

	if dbConnection == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, err := gorm.Open("mysql", dsn)

		//dsn := "root:admin@tcp(172.17.0.4:3306)/food-service?charset=utf8mb4&parseTime=True&loc=Local"
		//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			return nil, err
		}
		if err := db.DB().Ping(); err != nil {
			return nil, err
		}
		return db, nil

	} else if dbConnection == "postgres" {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
		db, err := gorm.Open("postgres", dsn)
		if err != nil {
			return nil, err
		}
		if err := db.DB().Ping(); err != nil {
			return nil, err
		}
		return db, nil

	} else {
		panic("invalid DB_CONNECTION (only mysql)")
	}
}
