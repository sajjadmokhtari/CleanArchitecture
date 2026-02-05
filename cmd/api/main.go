package main

import (
	"CleanArchitecture/internal/router"
	"log"
)

func main() {
	r := router.SetupRoutes()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
