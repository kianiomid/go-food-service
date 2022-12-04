package main

import (
	"fmt"
	"food-service/internal/presentation/route"
	"food-service/pkg/config"
	"food-service/pkg/database"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config.GetConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	//defer func(db *gorm.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(db)

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))

	app := route.SetupRouter(db)
	err = app.Run(port)
	if err != nil {
		log.Fatal(err)
		return
	}

}
