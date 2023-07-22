package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rals/dearme-channel/config"
	logs "github.com/rals/dearme-channel/logger"
	"github.com/rals/dearme-channel/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logs.InfoLogger.Println("- Info Main : No .env file found ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	config.InitializeDB()

	router := routes.InitializeRoute()
	router.Run()
}
