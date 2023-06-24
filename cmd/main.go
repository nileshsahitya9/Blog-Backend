package main

import (
	"log"

	"github.com/nileshsahitya9/Blogs-Backend/db"
	"github.com/nileshsahitya9/Blogs-Backend/internal/server"
)

func main() {
	database, err := db.PostgreSQLDB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer database.Close()

	s := server.NewServer()
	log.Fatal(s.Start(":8080"))

}
