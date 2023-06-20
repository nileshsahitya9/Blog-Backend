package main

import (
	"log"

	"github.com/nileshsahitya9/Blogs-Backend/internal/server"
)

func main() {
	s := server.NewServer()
	log.Fatal(s.Start(":8080"))
}
