package main

import (
	"article/config/database"
	"log"
)

func main() {
	dbConnection, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dbConnection)
}
