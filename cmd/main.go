package main

import (
	"article/config/database"
	interfaces "article/interfaces/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	interfaces.NewRouter(r, dbConnection)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println(dbConnection)
}
