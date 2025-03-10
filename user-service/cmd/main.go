package main

import (
	"fmt"
	"log"
	"user-service/config"
	"user-service/infrastructure"
)

func main() {
	// Inisialisasi database
	db := config.InitDB()

	// Setup router
	r := infrastructure.SetupRouter(db)

	// Start server
	port := ":8080"
	fmt.Println("User Service running on", port)
	log.Fatal(r.Run(port))
}
