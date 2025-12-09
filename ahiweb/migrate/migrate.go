package main

import (
	"ahiweb/ahiweb/database"
	"ahiweb/ahiweb/initializers"
	"ahiweb/ahiweb/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)
}

func main() {
	database.GlobalDB.AutoMigrate(&models.Users{})
	database.GlobalDB.AutoMigrate(&models.Posts{})
	database.GlobalDB.AutoMigrate(&models.Comments{})
	fmt.Println("? Migration complete")
}
