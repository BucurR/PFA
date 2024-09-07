package main

import (
	"log"
	"user-service/api"
	"user-service/config"
	"user-service/db"
)

func main() {
	conf := config.NewConfig()
	conf.InitConfig()
	db := db.NewAzureSQLDb(conf.Server, conf.Port, conf.Database, conf.User, conf.Passwrod)
	if err := db.ConnectDB(); err != nil {
		log.Println("Imi bag pl")
	}
	server := api.NewAPIServer(":8080", db.DB)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
