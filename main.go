package main

import (
	"database/sql"
	"fmt"
	"log"

	"myapp/controller"
	"myapp/migrations"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	err = migrations.CreateImagesTable(db)
	if err != nil {
		log.Fatalf("Error creating images table: %s", err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	controller.RegisterImageRoutes(router)

	router.Run(":8080")
}


