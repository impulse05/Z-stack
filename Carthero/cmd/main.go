package main

import (
	"Carthero/controllers"
	"Carthero/model"
	"Carthero/repository"
	"Carthero/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Rider{})

	riderRepo := repository.NewRiderRepository(db)
	riderService := services.NewRiderService(riderRepo)
	riderController := controllers.NewRiderController(riderService)

	r := gin.Default()

	r.POST("/riders", riderController.CreateRider)
	r.DELETE("/riders/:id", riderController.DeleteRider)
	r.GET("/riders/:id/:state", riderController.UpdateRiderState)
	r.GET("/riders/free", riderController.GetFreeRiders)

	r.Run(":8081")
}
