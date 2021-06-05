package main

import (
	"github.com/joho/godotenv"
	"log"
	"owen2020/app"

	_ "net/http/pprof"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.StartWebServer()
}
