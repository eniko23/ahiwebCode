package main

import (
	"ahiweb/ahiweb/database"
	"ahiweb/ahiweb/initializers"
	"ahiweb/ahiweb/router"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)
}

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
