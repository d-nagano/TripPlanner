package main

import (
	"fmt"
	"log"

	"trip-planner/db"
	"trip-planner/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}

	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func() {
		sqlDB, _ := dbConn.DB()
		if err := sqlDB.Close(); err != nil {
			log.Fatalln(err)
			return
		}
	}()

	if err := dbConn.AutoMigrate(&models.User{}, &models.TripPlan{},&models.TripDay{}); err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Successfully Migrated")
}
