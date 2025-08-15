package main

import (
	"backend/db"
	"backend/router"
	"log"
)

func main() {
	db.InitDB()

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
